// Copyright (c) 2017 Yandex LLC. All rights reserved.
// Use of this source code is governed by a MPL 2.0
// license that can be found in the LICENSE file.
// Author: Vladimir Skipor <skipor@yandex-team.ru>

package main

import (
	"github.com/yandex/pandora/core/aggregator/netsample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"time"

	_ "github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api"
	pb "github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api"
	log "github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"github.com/yandex/pandora/cli"
	phttp "github.com/yandex/pandora/components/phttp/import"
	"github.com/yandex/pandora/core"
	coreimport "github.com/yandex/pandora/core/import"
	"github.com/yandex/pandora/core/register"
)

type Ammo struct {
	Tag   string
	Param string
}

type Sample struct {
	Url              string
	ShootTimeSeconds float64
}

type GunConfig struct {
	Target string `validate:"required"` // Configuration will fail, without target defined
}

type Gun struct {
	// Configured on construction.
	conf GunConfig

	client grpc.ClientConn
	// Configured on Bind, before shooting.
	aggr core.Aggregator // May be your custom Aggregator.
	core.GunDeps
}

func NewGun(conf GunConfig) *Gun {
	return &Gun{conf: conf}
}

func (g *Gun) Bind(aggr core.Aggregator, deps core.GunDeps) error {
	conn, err := grpc.Dial(
		g.conf.Target,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
		grpc.WithUserAgent("load test, pandora custom shooter"))

	if err != nil {
		log.Error().Err(err).Msg("error bind")
	}
	g.client = *conn
	g.aggr = aggr
	g.GunDeps = deps
	return nil
}

func (g *Gun) Shoot(ammo core.Ammo) {
	customAmmo := ammo.(*Ammo) // Shoot will panic on unexpected ammo type. Panic cancels shooting.
	g.shoot(customAmmo)
}

func (g *Gun) shoot(ammo *Ammo) {
	start := time.Now()

	code := 0
	sample := netsample.Acquire(ammo.Tag)

	conn := g.client
	client := pb.NewOcpCertificateApiClient(&conn)

	// Put your shoot logic here.'
	req := http.Request{}
	out, err := client.CreateCertificateV1(req.Context(), &pb.CreateCertificateV1Request{
		Certificate: &pb.NewCertificate{
			UserId:    1,
			Created:   timestamppb.New(time.Now()),
			Link:      "https:link.ru",
			IsDeleted: false,
		},
	}, grpc.Header(&metadata.MD{}), grpc.Trailer(&metadata.MD{}))

	if err != nil {
		code = 0
	}

	if out != nil {
		code = 200
	}

	defer func() {
		sample.SetProtoCode(code)
		g.aggr.Report(Sample{ammo.Tag, time.Since(start).Seconds()})
	}()

}

func main() {
	// Standard imports.
	fs := afero.NewOsFs()
	coreimport.Import(fs)

	// May not be imported, if you don't need http guns and etc.
	phttp.Import(fs)

	// Custom imports. Integrate your custom types into configuration system.
	coreimport.RegisterCustomJSONProvider("ocp-certificate-api-provider", func() core.Ammo { return &Ammo{} })

	register.Gun("ocp-certificate-api-gun", NewGun, func() GunConfig {
		return GunConfig{
			Target: "default target",
		}
	})

	cli.Run()
}
