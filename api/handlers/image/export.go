// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package image

import (
	"net/http"

	"github.com/containerd/containerd/v2/pkg/namespaces"
	"github.com/gorilla/mux"

	"github.com/runfinch/finch-daemon/api/response"
	"github.com/runfinch/finch-daemon/pkg/errdefs"
)

func (h *handler) export(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	ctx := namespaces.WithNamespace(r.Context(), h.Config.Namespace)

	w.Header().Set("Content-Type", "application/x-tar")
	err := h.service.Export(ctx, name, w)
	if err != nil {
		var code int
		switch {
		case errdefs.IsNotFound(err):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}
		h.logger.Debugf("Export Image API failed. Status code %d, Message: %s", code, err)
		response.SendErrorResponse(w, code, err)
		return
	}
}
