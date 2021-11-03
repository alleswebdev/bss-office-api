# Generated by the Protocol Buffers compiler. DO NOT EDIT!
# source: ozonmp/bss_office_api/v1/bss_office_api.proto
# plugin: grpclib.plugin.main
import abc
import typing

import grpclib.const
import grpclib.client
if typing.TYPE_CHECKING:
    import grpclib.server

import validate.validate_pb2
import google.api.annotations_pb2
import google.protobuf.timestamp_pb2
import ozonmp.bss_office_api.v1.bss_office_api_pb2


class BssOfficeApiServiceBase(abc.ABC):

    @abc.abstractmethod
    async def DescribeOfficeV1(self, stream: 'grpclib.server.Stream[ozonmp.bss_office_api.v1.bss_office_api_pb2.DescribeOfficeV1Request, ozonmp.bss_office_api.v1.bss_office_api_pb2.DescribeOfficeV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def CreateOfficeV1(self, stream: 'grpclib.server.Stream[ozonmp.bss_office_api.v1.bss_office_api_pb2.CreateOfficeV1Request, ozonmp.bss_office_api.v1.bss_office_api_pb2.CreateOfficeV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def RemoveOfficeV1(self, stream: 'grpclib.server.Stream[ozonmp.bss_office_api.v1.bss_office_api_pb2.RemoveOfficeV1Request, ozonmp.bss_office_api.v1.bss_office_api_pb2.RemoveOfficeV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def ListOfficesV1(self, stream: 'grpclib.server.Stream[ozonmp.bss_office_api.v1.bss_office_api_pb2.ListOfficesV1Request, ozonmp.bss_office_api.v1.bss_office_api_pb2.ListOfficesV1Response]') -> None:
        pass

    def __mapping__(self) -> typing.Dict[str, grpclib.const.Handler]:
        return {
            '/ozonmp.bss_office_api.v1.BssOfficeApiService/DescribeOfficeV1': grpclib.const.Handler(
                self.DescribeOfficeV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.bss_office_api.v1.bss_office_api_pb2.DescribeOfficeV1Request,
                ozonmp.bss_office_api.v1.bss_office_api_pb2.DescribeOfficeV1Response,
            ),
            '/ozonmp.bss_office_api.v1.BssOfficeApiService/CreateOfficeV1': grpclib.const.Handler(
                self.CreateOfficeV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.bss_office_api.v1.bss_office_api_pb2.CreateOfficeV1Request,
                ozonmp.bss_office_api.v1.bss_office_api_pb2.CreateOfficeV1Response,
            ),
            '/ozonmp.bss_office_api.v1.BssOfficeApiService/RemoveOfficeV1': grpclib.const.Handler(
                self.RemoveOfficeV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.bss_office_api.v1.bss_office_api_pb2.RemoveOfficeV1Request,
                ozonmp.bss_office_api.v1.bss_office_api_pb2.RemoveOfficeV1Response,
            ),
            '/ozonmp.bss_office_api.v1.BssOfficeApiService/ListOfficesV1': grpclib.const.Handler(
                self.ListOfficesV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.bss_office_api.v1.bss_office_api_pb2.ListOfficesV1Request,
                ozonmp.bss_office_api.v1.bss_office_api_pb2.ListOfficesV1Response,
            ),
        }


class BssOfficeApiServiceStub:

    def __init__(self, channel: grpclib.client.Channel) -> None:
        self.DescribeOfficeV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.bss_office_api.v1.BssOfficeApiService/DescribeOfficeV1',
            ozonmp.bss_office_api.v1.bss_office_api_pb2.DescribeOfficeV1Request,
            ozonmp.bss_office_api.v1.bss_office_api_pb2.DescribeOfficeV1Response,
        )
        self.CreateOfficeV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.bss_office_api.v1.BssOfficeApiService/CreateOfficeV1',
            ozonmp.bss_office_api.v1.bss_office_api_pb2.CreateOfficeV1Request,
            ozonmp.bss_office_api.v1.bss_office_api_pb2.CreateOfficeV1Response,
        )
        self.RemoveOfficeV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.bss_office_api.v1.BssOfficeApiService/RemoveOfficeV1',
            ozonmp.bss_office_api.v1.bss_office_api_pb2.RemoveOfficeV1Request,
            ozonmp.bss_office_api.v1.bss_office_api_pb2.RemoveOfficeV1Response,
        )
        self.ListOfficesV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.bss_office_api.v1.BssOfficeApiService/ListOfficesV1',
            ozonmp.bss_office_api.v1.bss_office_api_pb2.ListOfficesV1Request,
            ozonmp.bss_office_api.v1.bss_office_api_pb2.ListOfficesV1Response,
        )