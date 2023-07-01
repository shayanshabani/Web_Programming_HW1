import interceptor
import AuthService_pb2
import AuthService_pb2_grpc
from locust import task


class AuthGrpcUser(interceptor.GrpcUser):
    host = "localhost:5052"
    stub_class = AuthService_pb2_grpc.Auth_ServiceStub
    server_nonce = ""

    @task
    def reqPq(self):
        self.stub.req_pq(AuthService_pb2.req_pq_Request(nonce="abcd", message_id=0))
