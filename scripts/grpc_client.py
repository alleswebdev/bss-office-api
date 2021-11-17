import asyncio

from grpclib.client import Channel

from ozonmp.bss_office_api.v1.bss_office_api_grpc import BssOfficeApiServiceStub
from ozonmp.bss_office_api.v1.bss_office_api_pb2 import DescribeOfficeV1Request


async def main():
    async with Channel('127.0.0.1', 8082) as channel:
        client = BssOfficeApiServiceStub(channel)

        req = DescribeOfficeV1Request(office_id=1)
        reply = await client.DescribeOfficeV1(req)
        print(reply)


if __name__ == '__main__':
    asyncio.run(main())
