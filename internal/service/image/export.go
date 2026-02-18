// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package image

import (
	"context"
	"io"

	"github.com/containerd/containerd/v2/core/images"
)

func (s *service) Export(ctx context.Context, name string, outStream io.Writer) error {
	img, err := s.getImage(ctx, name)
	if err != nil {
		return err
	}
	return s.nctlImageSvc.ExportImage(ctx, []images.Image{*img}, outStream)
}
