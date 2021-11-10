# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ozonmp/bss_office_api/v1/bss_office_api.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from validate import validate_pb2 as validate_dot_validate__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='ozonmp/bss_office_api/v1/bss_office_api.proto',
  package='ozonmp.bss_office_api.v1',
  syntax='proto3',
  serialized_options=b'ZBgithub.com/ozonmp/bss-office-api/pkg/bss-office-api;bss_office_api',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n-ozonmp/bss_office_api/v1/bss_office_api.proto\x12\x18ozonmp.bss_office_api.v1\x1a\x17validate/validate.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xe8\x01\n\x06Office\x12\x17\n\x02id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x02id\x12\x1d\n\x04name\x18\x02 \x01(\tB\t\xfa\x42\x06r\x04\x10\x02\x18\x64R\x04name\x12 \n\x0b\x64\x65scription\x18\x03 \x01(\tR\x0b\x64\x65scription\x12\x18\n\x07removed\x18\x04 \x01(\x08R\x07removed\x12\x34\n\x07\x63reated\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x07\x63reated\x12\x34\n\x07updated\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x07updated\"\xf5\x01\n\x0bOfficeEvent\x12\x17\n\x02id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x02id\x12$\n\toffice_id\x18\x02 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x08officeId\x12\x16\n\x06status\x18\x03 \x01(\x04R\x06status\x12\x1d\n\x04type\x18\x04 \x01(\tB\t\xfa\x42\x06r\x04\x10\x02\x18\x64R\x04type\x12\x34\n\x07\x63reated\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x07\x63reated\x12:\n\x07payload\x18\x06 \x01(\x0b\x32 .ozonmp.bss_office_api.v1.OfficeR\x07payload\"?\n\x17\x44\x65scribeOfficeV1Request\x12$\n\toffice_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x08officeId\"R\n\x18\x44\x65scribeOfficeV1Response\x12\x36\n\x05value\x18\x01 \x01(\x0b\x32 .ozonmp.bss_office_api.v1.OfficeR\x05value\"X\n\x15\x43reateOfficeV1Request\x12\x1d\n\x04name\x18\x01 \x01(\tB\t\xfa\x42\x06r\x04\x10\x02\x18\x64R\x04name\x12 \n\x0b\x64\x65scription\x18\x02 \x01(\tR\x0b\x64\x65scription\">\n\x16\x43reateOfficeV1Response\x12$\n\toffice_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x08officeId\"O\n\x14ListOfficesV1Request\x12\x1f\n\x05limit\x18\x01 \x01(\x04\x42\t\xfa\x42\x06\x32\x04\x10\x64 \x00R\x05limit\x12\x16\n\x06offset\x18\x02 \x01(\x04R\x06offset\"O\n\x15ListOfficesV1Response\x12\x36\n\x05items\x18\x01 \x03(\x0b\x32 .ozonmp.bss_office_api.v1.OfficeR\x05items\"=\n\x15RemoveOfficeV1Request\x12$\n\toffice_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x08officeId\".\n\x16RemoveOfficeV1Response\x12\x14\n\x05\x66ound\x18\x01 \x01(\x08R\x05\x66ound\"~\n\x15UpdateOfficeV1Request\x12$\n\toffice_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x08officeId\x12\x1d\n\x04name\x18\x02 \x01(\tB\t\xfa\x42\x06r\x04\x10\x02\x18\x64R\x04name\x12 \n\x0b\x64\x65scription\x18\x03 \x01(\tR\x0b\x64\x65scription\"0\n\x16UpdateOfficeV1Response\x12\x16\n\x06status\x18\x01 \x01(\x08R\x06status2\x92\x06\n\x13\x42ssOfficeApiService\x12\x9d\x01\n\x10\x44\x65scribeOfficeV1\x12\x31.ozonmp.bss_office_api.v1.DescribeOfficeV1Request\x1a\x32.ozonmp.bss_office_api.v1.DescribeOfficeV1Response\"\"\x82\xd3\xe4\x93\x02\x1c\x12\x1a/api/v1/office/{office_id}\x12\x8e\x01\n\x0e\x43reateOfficeV1\x12/.ozonmp.bss_office_api.v1.CreateOfficeV1Request\x1a\x30.ozonmp.bss_office_api.v1.CreateOfficeV1Response\"\x19\x82\xd3\xe4\x93\x02\x13\"\x0e/api/v1/office:\x01*\x12\x97\x01\n\x0eRemoveOfficeV1\x12/.ozonmp.bss_office_api.v1.RemoveOfficeV1Request\x1a\x30.ozonmp.bss_office_api.v1.RemoveOfficeV1Response\"\"\x82\xd3\xe4\x93\x02\x1c*\x1a/api/v1/office/{office_id}\x12\x9e\x01\n\rListOfficesV1\x12..ozonmp.bss_office_api.v1.ListOfficesV1Request\x1a/.ozonmp.bss_office_api.v1.ListOfficesV1Response\",\x82\xd3\xe4\x93\x02&\x12$/api/v1/office/list/{limit}/{offset}\x12\x8e\x01\n\x0eUpdateOfficeV1\x12/.ozonmp.bss_office_api.v1.UpdateOfficeV1Request\x1a\x30.ozonmp.bss_office_api.v1.UpdateOfficeV1Response\"\x19\x82\xd3\xe4\x93\x02\x13\x1a\x0e/api/v1/office:\x01*BDZBgithub.com/ozonmp/bss-office-api/pkg/bss-office-api;bss_office_apib\x06proto3'
  ,
  dependencies=[validate_dot_validate__pb2.DESCRIPTOR,google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_OFFICE = _descriptor.Descriptor(
  name='Office',
  full_name='ozonmp.bss_office_api.v1.Office',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='ozonmp.bss_office_api.v1.Office.id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\0042\002 \000', json_name='id', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='ozonmp.bss_office_api.v1.Office.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\006r\004\020\002\030d', json_name='name', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='description', full_name='ozonmp.bss_office_api.v1.Office.description', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='description', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='removed', full_name='ozonmp.bss_office_api.v1.Office.removed', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='removed', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='created', full_name='ozonmp.bss_office_api.v1.Office.created', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='created', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='updated', full_name='ozonmp.bss_office_api.v1.Office.updated', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='updated', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=164,
  serialized_end=396,
)


_OFFICEEVENT = _descriptor.Descriptor(
  name='OfficeEvent',
  full_name='ozonmp.bss_office_api.v1.OfficeEvent',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='ozonmp.bss_office_api.v1.OfficeEvent.id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\0042\002 \000', json_name='id', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='office_id', full_name='ozonmp.bss_office_api.v1.OfficeEvent.office_id', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\0042\002 \000', json_name='officeId', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='status', full_name='ozonmp.bss_office_api.v1.OfficeEvent.status', index=2,
      number=3, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='status', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='type', full_name='ozonmp.bss_office_api.v1.OfficeEvent.type', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\006r\004\020\002\030d', json_name='type', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='created', full_name='ozonmp.bss_office_api.v1.OfficeEvent.created', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='created', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='payload', full_name='ozonmp.bss_office_api.v1.OfficeEvent.payload', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='payload', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=399,
  serialized_end=644,
)


_DESCRIBEOFFICEV1REQUEST = _descriptor.Descriptor(
  name='DescribeOfficeV1Request',
  full_name='ozonmp.bss_office_api.v1.DescribeOfficeV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='office_id', full_name='ozonmp.bss_office_api.v1.DescribeOfficeV1Request.office_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\0042\002 \000', json_name='officeId', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=646,
  serialized_end=709,
)


_DESCRIBEOFFICEV1RESPONSE = _descriptor.Descriptor(
  name='DescribeOfficeV1Response',
  full_name='ozonmp.bss_office_api.v1.DescribeOfficeV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='ozonmp.bss_office_api.v1.DescribeOfficeV1Response.value', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='value', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=711,
  serialized_end=793,
)


_CREATEOFFICEV1REQUEST = _descriptor.Descriptor(
  name='CreateOfficeV1Request',
  full_name='ozonmp.bss_office_api.v1.CreateOfficeV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='ozonmp.bss_office_api.v1.CreateOfficeV1Request.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\006r\004\020\002\030d', json_name='name', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='description', full_name='ozonmp.bss_office_api.v1.CreateOfficeV1Request.description', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='description', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=795,
  serialized_end=883,
)


_CREATEOFFICEV1RESPONSE = _descriptor.Descriptor(
  name='CreateOfficeV1Response',
  full_name='ozonmp.bss_office_api.v1.CreateOfficeV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='office_id', full_name='ozonmp.bss_office_api.v1.CreateOfficeV1Response.office_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\0042\002 \000', json_name='officeId', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=885,
  serialized_end=947,
)


_LISTOFFICESV1REQUEST = _descriptor.Descriptor(
  name='ListOfficesV1Request',
  full_name='ozonmp.bss_office_api.v1.ListOfficesV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='limit', full_name='ozonmp.bss_office_api.v1.ListOfficesV1Request.limit', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\0062\004\020d \000', json_name='limit', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='offset', full_name='ozonmp.bss_office_api.v1.ListOfficesV1Request.offset', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='offset', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=949,
  serialized_end=1028,
)


_LISTOFFICESV1RESPONSE = _descriptor.Descriptor(
  name='ListOfficesV1Response',
  full_name='ozonmp.bss_office_api.v1.ListOfficesV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='items', full_name='ozonmp.bss_office_api.v1.ListOfficesV1Response.items', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='items', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1030,
  serialized_end=1109,
)


_REMOVEOFFICEV1REQUEST = _descriptor.Descriptor(
  name='RemoveOfficeV1Request',
  full_name='ozonmp.bss_office_api.v1.RemoveOfficeV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='office_id', full_name='ozonmp.bss_office_api.v1.RemoveOfficeV1Request.office_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\0042\002 \000', json_name='officeId', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1111,
  serialized_end=1172,
)


_REMOVEOFFICEV1RESPONSE = _descriptor.Descriptor(
  name='RemoveOfficeV1Response',
  full_name='ozonmp.bss_office_api.v1.RemoveOfficeV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='found', full_name='ozonmp.bss_office_api.v1.RemoveOfficeV1Response.found', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='found', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1174,
  serialized_end=1220,
)


_UPDATEOFFICEV1REQUEST = _descriptor.Descriptor(
  name='UpdateOfficeV1Request',
  full_name='ozonmp.bss_office_api.v1.UpdateOfficeV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='office_id', full_name='ozonmp.bss_office_api.v1.UpdateOfficeV1Request.office_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\0042\002 \000', json_name='officeId', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='ozonmp.bss_office_api.v1.UpdateOfficeV1Request.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\006r\004\020\002\030d', json_name='name', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='description', full_name='ozonmp.bss_office_api.v1.UpdateOfficeV1Request.description', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='description', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1222,
  serialized_end=1348,
)


_UPDATEOFFICEV1RESPONSE = _descriptor.Descriptor(
  name='UpdateOfficeV1Response',
  full_name='ozonmp.bss_office_api.v1.UpdateOfficeV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='ozonmp.bss_office_api.v1.UpdateOfficeV1Response.status', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='status', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1350,
  serialized_end=1398,
)

_OFFICE.fields_by_name['created'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_OFFICE.fields_by_name['updated'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_OFFICEEVENT.fields_by_name['created'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_OFFICEEVENT.fields_by_name['payload'].message_type = _OFFICE
_DESCRIBEOFFICEV1RESPONSE.fields_by_name['value'].message_type = _OFFICE
_LISTOFFICESV1RESPONSE.fields_by_name['items'].message_type = _OFFICE
DESCRIPTOR.message_types_by_name['Office'] = _OFFICE
DESCRIPTOR.message_types_by_name['OfficeEvent'] = _OFFICEEVENT
DESCRIPTOR.message_types_by_name['DescribeOfficeV1Request'] = _DESCRIBEOFFICEV1REQUEST
DESCRIPTOR.message_types_by_name['DescribeOfficeV1Response'] = _DESCRIBEOFFICEV1RESPONSE
DESCRIPTOR.message_types_by_name['CreateOfficeV1Request'] = _CREATEOFFICEV1REQUEST
DESCRIPTOR.message_types_by_name['CreateOfficeV1Response'] = _CREATEOFFICEV1RESPONSE
DESCRIPTOR.message_types_by_name['ListOfficesV1Request'] = _LISTOFFICESV1REQUEST
DESCRIPTOR.message_types_by_name['ListOfficesV1Response'] = _LISTOFFICESV1RESPONSE
DESCRIPTOR.message_types_by_name['RemoveOfficeV1Request'] = _REMOVEOFFICEV1REQUEST
DESCRIPTOR.message_types_by_name['RemoveOfficeV1Response'] = _REMOVEOFFICEV1RESPONSE
DESCRIPTOR.message_types_by_name['UpdateOfficeV1Request'] = _UPDATEOFFICEV1REQUEST
DESCRIPTOR.message_types_by_name['UpdateOfficeV1Response'] = _UPDATEOFFICEV1RESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Office = _reflection.GeneratedProtocolMessageType('Office', (_message.Message,), {
  'DESCRIPTOR' : _OFFICE,
  '__module__' : 'ozonmp.bss_office_api.v1.bss_office_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_office_api.v1.Office)
  })
_sym_db.RegisterMessage(Office)

OfficeEvent = _reflection.GeneratedProtocolMessageType('OfficeEvent', (_message.Message,), {
  'DESCRIPTOR' : _OFFICEEVENT,
  '__module__' : 'ozonmp.bss_office_api.v1.bss_office_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_office_api.v1.OfficeEvent)
  })
_sym_db.RegisterMessage(OfficeEvent)

DescribeOfficeV1Request = _reflection.GeneratedProtocolMessageType('DescribeOfficeV1Request', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBEOFFICEV1REQUEST,
  '__module__' : 'ozonmp.bss_office_api.v1.bss_office_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_office_api.v1.DescribeOfficeV1Request)
  })
_sym_db.RegisterMessage(DescribeOfficeV1Request)

DescribeOfficeV1Response = _reflection.GeneratedProtocolMessageType('DescribeOfficeV1Response', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBEOFFICEV1RESPONSE,
  '__module__' : 'ozonmp.bss_office_api.v1.bss_office_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_office_api.v1.DescribeOfficeV1Response)
  })
_sym_db.RegisterMessage(DescribeOfficeV1Response)

CreateOfficeV1Request = _reflection.GeneratedProtocolMessageType('CreateOfficeV1Request', (_message.Message,), {
  'DESCRIPTOR' : _CREATEOFFICEV1REQUEST,
  '__module__' : 'ozonmp.bss_office_api.v1.bss_office_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_office_api.v1.CreateOfficeV1Request)
  })
_sym_db.RegisterMessage(CreateOfficeV1Request)

CreateOfficeV1Response = _reflection.GeneratedProtocolMessageType('CreateOfficeV1Response', (_message.Message,), {
  'DESCRIPTOR' : _CREATEOFFICEV1RESPONSE,
  '__module__' : 'ozonmp.bss_office_api.v1.bss_office_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_office_api.v1.CreateOfficeV1Response)
  })
_sym_db.RegisterMessage(CreateOfficeV1Response)

ListOfficesV1Request = _reflection.GeneratedProtocolMessageType('ListOfficesV1Request', (_message.Message,), {
  'DESCRIPTOR' : _LISTOFFICESV1REQUEST,
  '__module__' : 'ozonmp.bss_office_api.v1.bss_office_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_office_api.v1.ListOfficesV1Request)
  })
_sym_db.RegisterMessage(ListOfficesV1Request)

ListOfficesV1Response = _reflection.GeneratedProtocolMessageType('ListOfficesV1Response', (_message.Message,), {
  'DESCRIPTOR' : _LISTOFFICESV1RESPONSE,
  '__module__' : 'ozonmp.bss_office_api.v1.bss_office_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_office_api.v1.ListOfficesV1Response)
  })
_sym_db.RegisterMessage(ListOfficesV1Response)

RemoveOfficeV1Request = _reflection.GeneratedProtocolMessageType('RemoveOfficeV1Request', (_message.Message,), {
  'DESCRIPTOR' : _REMOVEOFFICEV1REQUEST,
  '__module__' : 'ozonmp.bss_office_api.v1.bss_office_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_office_api.v1.RemoveOfficeV1Request)
  })
_sym_db.RegisterMessage(RemoveOfficeV1Request)

RemoveOfficeV1Response = _reflection.GeneratedProtocolMessageType('RemoveOfficeV1Response', (_message.Message,), {
  'DESCRIPTOR' : _REMOVEOFFICEV1RESPONSE,
  '__module__' : 'ozonmp.bss_office_api.v1.bss_office_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_office_api.v1.RemoveOfficeV1Response)
  })
_sym_db.RegisterMessage(RemoveOfficeV1Response)

UpdateOfficeV1Request = _reflection.GeneratedProtocolMessageType('UpdateOfficeV1Request', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEOFFICEV1REQUEST,
  '__module__' : 'ozonmp.bss_office_api.v1.bss_office_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_office_api.v1.UpdateOfficeV1Request)
  })
_sym_db.RegisterMessage(UpdateOfficeV1Request)

UpdateOfficeV1Response = _reflection.GeneratedProtocolMessageType('UpdateOfficeV1Response', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEOFFICEV1RESPONSE,
  '__module__' : 'ozonmp.bss_office_api.v1.bss_office_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_office_api.v1.UpdateOfficeV1Response)
  })
_sym_db.RegisterMessage(UpdateOfficeV1Response)


DESCRIPTOR._options = None
_OFFICE.fields_by_name['id']._options = None
_OFFICE.fields_by_name['name']._options = None
_OFFICEEVENT.fields_by_name['id']._options = None
_OFFICEEVENT.fields_by_name['office_id']._options = None
_OFFICEEVENT.fields_by_name['type']._options = None
_DESCRIBEOFFICEV1REQUEST.fields_by_name['office_id']._options = None
_CREATEOFFICEV1REQUEST.fields_by_name['name']._options = None
_CREATEOFFICEV1RESPONSE.fields_by_name['office_id']._options = None
_LISTOFFICESV1REQUEST.fields_by_name['limit']._options = None
_REMOVEOFFICEV1REQUEST.fields_by_name['office_id']._options = None
_UPDATEOFFICEV1REQUEST.fields_by_name['office_id']._options = None
_UPDATEOFFICEV1REQUEST.fields_by_name['name']._options = None

_BSSOFFICEAPISERVICE = _descriptor.ServiceDescriptor(
  name='BssOfficeApiService',
  full_name='ozonmp.bss_office_api.v1.BssOfficeApiService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=1401,
  serialized_end=2187,
  methods=[
  _descriptor.MethodDescriptor(
    name='DescribeOfficeV1',
    full_name='ozonmp.bss_office_api.v1.BssOfficeApiService.DescribeOfficeV1',
    index=0,
    containing_service=None,
    input_type=_DESCRIBEOFFICEV1REQUEST,
    output_type=_DESCRIBEOFFICEV1RESPONSE,
    serialized_options=b'\202\323\344\223\002\034\022\032/api/v1/office/{office_id}',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='CreateOfficeV1',
    full_name='ozonmp.bss_office_api.v1.BssOfficeApiService.CreateOfficeV1',
    index=1,
    containing_service=None,
    input_type=_CREATEOFFICEV1REQUEST,
    output_type=_CREATEOFFICEV1RESPONSE,
    serialized_options=b'\202\323\344\223\002\023\"\016/api/v1/office:\001*',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='RemoveOfficeV1',
    full_name='ozonmp.bss_office_api.v1.BssOfficeApiService.RemoveOfficeV1',
    index=2,
    containing_service=None,
    input_type=_REMOVEOFFICEV1REQUEST,
    output_type=_REMOVEOFFICEV1RESPONSE,
    serialized_options=b'\202\323\344\223\002\034*\032/api/v1/office/{office_id}',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='ListOfficesV1',
    full_name='ozonmp.bss_office_api.v1.BssOfficeApiService.ListOfficesV1',
    index=3,
    containing_service=None,
    input_type=_LISTOFFICESV1REQUEST,
    output_type=_LISTOFFICESV1RESPONSE,
    serialized_options=b'\202\323\344\223\002&\022$/api/v1/office/list/{limit}/{offset}',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='UpdateOfficeV1',
    full_name='ozonmp.bss_office_api.v1.BssOfficeApiService.UpdateOfficeV1',
    index=4,
    containing_service=None,
    input_type=_UPDATEOFFICEV1REQUEST,
    output_type=_UPDATEOFFICEV1RESPONSE,
    serialized_options=b'\202\323\344\223\002\023\032\016/api/v1/office:\001*',
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_BSSOFFICEAPISERVICE)

DESCRIPTOR.services_by_name['BssOfficeApiService'] = _BSSOFFICEAPISERVICE

# @@protoc_insertion_point(module_scope)
