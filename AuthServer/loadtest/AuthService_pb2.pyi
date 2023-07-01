from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Confirm(_message.Message):
    __slots__ = ["valid"]
    VALID_FIELD_NUMBER: _ClassVar[int]
    valid: bool
    def __init__(self, valid: bool = ...) -> None: ...

class Seek(_message.Message):
    __slots__ = ["auth_key"]
    AUTH_KEY_FIELD_NUMBER: _ClassVar[int]
    auth_key: int
    def __init__(self, auth_key: _Optional[int] = ...) -> None: ...

class req_dh_Request(_message.Message):
    __slots__ = ["A", "message_id", "nonce", "server_nonce"]
    A: int
    A_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_ID_FIELD_NUMBER: _ClassVar[int]
    NONCE_FIELD_NUMBER: _ClassVar[int]
    SERVER_NONCE_FIELD_NUMBER: _ClassVar[int]
    message_id: int
    nonce: str
    server_nonce: str
    def __init__(self, nonce: _Optional[str] = ..., server_nonce: _Optional[str] = ..., message_id: _Optional[int] = ..., A: _Optional[int] = ...) -> None: ...

class req_dh_Response(_message.Message):
    __slots__ = ["B", "message_id", "nonce", "server_nonce"]
    B: int
    B_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_ID_FIELD_NUMBER: _ClassVar[int]
    NONCE_FIELD_NUMBER: _ClassVar[int]
    SERVER_NONCE_FIELD_NUMBER: _ClassVar[int]
    message_id: int
    nonce: str
    server_nonce: str
    def __init__(self, nonce: _Optional[str] = ..., server_nonce: _Optional[str] = ..., message_id: _Optional[int] = ..., B: _Optional[int] = ...) -> None: ...

class req_pq_Request(_message.Message):
    __slots__ = ["message_id", "nonce"]
    MESSAGE_ID_FIELD_NUMBER: _ClassVar[int]
    NONCE_FIELD_NUMBER: _ClassVar[int]
    message_id: int
    nonce: str
    def __init__(self, nonce: _Optional[str] = ..., message_id: _Optional[int] = ...) -> None: ...

class req_pq_Response(_message.Message):
    __slots__ = ["g", "message_id", "nonce", "p", "server_nonce"]
    G_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_ID_FIELD_NUMBER: _ClassVar[int]
    NONCE_FIELD_NUMBER: _ClassVar[int]
    P_FIELD_NUMBER: _ClassVar[int]
    SERVER_NONCE_FIELD_NUMBER: _ClassVar[int]
    g: int
    message_id: int
    nonce: str
    p: int
    server_nonce: str
    def __init__(self, nonce: _Optional[str] = ..., server_nonce: _Optional[str] = ..., message_id: _Optional[int] = ..., p: _Optional[int] = ..., g: _Optional[int] = ...) -> None: ...
