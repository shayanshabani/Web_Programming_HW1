# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: BizServer.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0f\x42izServer.proto\x12\tBizServer\"^\n\x04User\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0e\n\x06\x66\x61mily\x18\x03 \x01(\t\x12\x0b\n\x03\x61ge\x18\x04 \x01(\x05\x12\x0b\n\x03sex\x18\x05 \x01(\t\x12\x12\n\ncreated_at\x18\x06 \x01(\t\"E\n\x0fGetUsersRequest\x12\x0e\n\x06userId\x18\x01 \x01(\x05\x12\x0f\n\x07\x61uthKey\x18\x02 \x01(\x05\x12\x11\n\tmessageId\x18\x03 \x01(\x05\"E\n\x10GetUsersResponse\x12\x1e\n\x05users\x18\x01 \x03(\x0b\x32\x0f.BizServer.User\x12\x11\n\tmessageId\x18\x02 \x01(\x05\"R\n\x1cGetUsersWithInjectionRequest\x12\x0e\n\x06userId\x18\x01 \x01(\t\x12\x0f\n\x07\x61uthKey\x18\x02 \x01(\x05\x12\x11\n\tmessageId\x18\x03 \x01(\x05\x32\xb9\x01\n\x0fGetUsersService\x12\x45\n\x08GetUsers\x12\x1a.BizServer.GetUsersRequest\x1a\x1b.BizServer.GetUsersResponse\"\x00\x12_\n\x15GetUsersWithInjection\x12\'.BizServer.GetUsersWithInjectionRequest\x1a\x1b.BizServer.GetUsersResponse\"\x00\x42\x0fZ\rweb_hw1/protob\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'BizServer_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\rweb_hw1/proto'
  _USER._serialized_start=30
  _USER._serialized_end=124
  _GETUSERSREQUEST._serialized_start=126
  _GETUSERSREQUEST._serialized_end=195
  _GETUSERSRESPONSE._serialized_start=197
  _GETUSERSRESPONSE._serialized_end=266
  _GETUSERSWITHINJECTIONREQUEST._serialized_start=268
  _GETUSERSWITHINJECTIONREQUEST._serialized_end=350
  _GETUSERSSERVICE._serialized_start=353
  _GETUSERSSERVICE._serialized_end=538
# @@protoc_insertion_point(module_scope)