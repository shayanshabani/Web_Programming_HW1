from locust import HttpUser, task, between


class WebsiteTestUser(HttpUser):
    wait_time = between(0.5, 3.0)

    @task
    def getUsers(self):
        self.client.post("/get-user", json={"userId": 1, "messageId": 1, "authKey": 1})

    @task
    def getUsersInject(self):
        self.client.post("/get-user-inject", json={"userId": "1", "messageId": 1, "authKey": 1})

    @task
    def reqPq(self):
        self.client.post("/auth/req_pq", json={"nonce": "abcd", "message_id": 0})
