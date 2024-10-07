package job

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetRunbookContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *string
}

type GetRunbookContentOperationOptions struct {
	ClientRequestId *string
}

func DefaultGetRunbookContentOperationOptions() GetRunbookContentOperationOptions {
	return GetRunbookContentOperationOptions{}
}

func (o GetRunbookContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.ClientRequestId != nil {
		out.Append("clientRequestId", fmt.Sprintf("%v", *o.ClientRequestId))
	}
	return &out
}

func (o GetRunbookContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o GetRunbookContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetRunbookContent ...
func (c JobClient) GetRunbookContent(ctx context.Context, id JobId, options GetRunbookContentOperationOptions) (result GetRunbookContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/powershell",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/runbookContent", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model string
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
