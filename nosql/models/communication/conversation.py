from nosql.models.model import Model
from pydantic import Field, HttpUrl
from typing import List, Optional
from bson import ObjectId
from ..persons.user import UserReadOnly
from nosql.models.communication.message import Message


class Conversation(Model):
    title: str = Field(..., max_length=32)
    icon: Optional[HttpUrl] = None
    author_id: ObjectId
    members: Optional[List[UserReadOnly]] = None
    message: Optional[List[Message]] = None
