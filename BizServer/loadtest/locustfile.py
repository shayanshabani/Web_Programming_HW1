import interceptor
import BizServer_pb2
import BizServer_pb2_grpc
from locust import task


class GetUsersGrpcUser(interceptor.GrpcUser):
    host = "localhost:5062"
    stub_class = BizServer_pb2_grpc.GetUsersServiceStub

    @task
    def getUsers(self):
        self.stub.GetUsers(BizServer_pb2.GetUsersRequest(userId=1, messageId=2, authKey=1))

    @task
    def getUsersWithInject(self):
        self.stub.GetUsersWithInjection(BizServer_pb2.GetUsersWithInjectionRequest(userId="1", messageId=2, authKey=1))
