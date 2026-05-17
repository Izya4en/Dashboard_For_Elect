package pb

import (
	context "context"

	grpc "google.golang.org/grpc"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat    float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon    float64 `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"`
	Weight float64 `protobuf:"fixed64,3,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return protoimpl.X.MessageOf(x)
}

type HeatmapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points []*Point `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
}

func (x *HeatmapRequest) Reset() {
	*x = HeatmapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeatmapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeatmapRequest) ProtoMessage() {}

func (x *HeatmapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return protoimpl.X.MessageOf(x)
}

type HeatmapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DensityPoints []*Point `protobuf:"bytes,1,rep,name=density_points,json=densityPoints,proto3" json:"density_points,omitempty"`
}

func (x *HeatmapResponse) Reset() {
	*x = HeatmapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeatmapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeatmapResponse) ProtoMessage() {}

func (x *HeatmapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return protoimpl.X.MessageOf(x)
}

type SpatialAnalyticsClient interface {
	CalculateHeatmap(ctx context.Context, in *HeatmapRequest, opts ...grpc.CallOption) (*HeatmapResponse, error)
}

type spatialAnalyticsClient struct {
	cc grpc.ClientConnInterface
}

func NewSpatialAnalyticsClient(cc grpc.ClientConnInterface) SpatialAnalyticsClient {
	return &spatialAnalyticsClient{cc}
}

func (c *spatialAnalyticsClient) CalculateHeatmap(ctx context.Context, in *HeatmapRequest, opts ...grpc.CallOption) (*HeatmapResponse, error) {
	out := new(HeatmapResponse)
	err := c.cc.Invoke(ctx, "/analytics.SpatialAnalytics/CalculateHeatmap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var file_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 3)

func init() {
	file_analytics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
		switch v := v.(*Point); i {
		case 0:
			return &v.state
		case 1:
			return &v.sizeCache
		case 2:
			return &v.unknownFields
		default:
			return nil
		}
	}
	file_analytics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
		switch v := v.(*HeatmapRequest); i {
		case 0:
			return &v.state
		case 1:
			return &v.sizeCache
		case 2:
			return &v.unknownFields
		default:
			return nil
		}
	}
	file_analytics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
		switch v := v.(*HeatmapResponse); i {
		case 0:
			return &v.state
		case 1:
			return &v.sizeCache
		case 2:
			return &v.unknownFields
		default:
			return nil
		}
	}
}
