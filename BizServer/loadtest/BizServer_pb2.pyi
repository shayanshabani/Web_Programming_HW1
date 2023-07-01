from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetUsersRequest(_message.Message):
    __slots__ = ["authKey", "messageId", "userId"]
    AUTHKEY_FIELD_NUMBER: _ClassVar[int]
    MESSAGEID_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    authKey: int
    messageId: int
    userId: int
    def __init__(self, userId: _Optional[int] = ..., authKey: _Optional[int] = ..., messageId: _Optional[int] = ...) -> None: ...

class GetUsersResponse(_message.Message):
    __slots__ = ["messageId", "users"]
    MESSAGEID_FIELD_NUMBER: _ClassVar[int]
    USERS_FIELD_NUMBER: _ClassVar[int]
    messageId: int
    users: _containers.RepeatedCompositeFieldContainer[User]
    def __init__(self, users: _Optional[_Iterable[_Union[User, _Mapping]]] = ..., messageId: _Optional[int] = ...) -> None: ...

class GetUsersWithInjectionRequest(_message.Message):
    __slots__ = ["authKey", "messageId", "userId"]
    AUTHKEY_FIELD_NUMBER: _ClassVar[int]
    MESSAGEID_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    authKey: int
    messageId: int
    userId: str
    def __init__(self, userId: _Optional[str] = ..., authKey: _Optional[int] = ..., messageId: _Optional[int] = ...) -> None: ...

class User(_message.Message):
    __slots__ = ["age", "created_at", "family", "id", "name", "sex"]
    AGE_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    FAMILY_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    SEX_FIELD_NUMBER: _ClassVar[int]
    age: int
    created_at: str
    family: str
    id: str
    name: str
    sex: str
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., family: _Optional[str] = ..., age: _Optional[int] = ..., sex: _Optional[str] = ..., created_at: _Optional[str] = ...) -> None: ...
