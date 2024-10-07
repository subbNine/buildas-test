def test_register_user():
    response = client.post("/register", json={"email": "test@example.com", "password": "test123"})
    assert response.status_code == 200
    assert response.json()["email"] == "test@example.com"

def test_list_users():
    response = client.get("/users")
    assert response.status_code == 200
    assert len(response.json()) > 0
