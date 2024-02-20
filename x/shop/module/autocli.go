package shop

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "shop/api/shop/shop"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "ShowPost",
					Use:            "show-post [id]",
					Short:          "Query show-post",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListPost",
					Use:            "list-post [start-date] [end-date] [min-price] [max-price]",
					Short:          "Query list-post",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "startDate"}, {ProtoField: "endDate"}, {ProtoField: "minPrice"}, {ProtoField: "maxPrice"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreatePost",
					Use:            "create-post [title] [description] [price]",
					Short:          "Send a create-post tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "price"}},
				},
				{
					RpcMethod:      "UpdatePost",
					Use:            "update-post [title] [description] [price] [id]",
					Short:          "Send a update-post tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "price"}, {ProtoField: "id"}},
				},
				{
					RpcMethod:      "DeletePost",
					Use:            "delete-post [id]",
					Short:          "Send a delete-post tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "ChangePostStatus",
					Use:            "change-post-status [id] [status]",
					Short:          "Send a change-post-status tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "status"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
