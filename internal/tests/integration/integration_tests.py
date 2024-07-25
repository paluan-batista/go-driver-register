import unittest
import requests

BASE_URL = "http://localhost:8000"


class TestDriverAPI(unittest.TestCase):
    def setUp(self):
        self.driver_id = None

    def test_create_driver(self):
        payload = {
            "name": "John Doe",
            "license_id": "12345"
        }
        response = requests.post(f"{BASE_URL}/drivers", json=payload)
        self.assertEqual(response.status_code, 201)
        data = response.json()
        self.assertIn("id", data)
        self.assertIsInstance(data["id"], str)
        self.driver_id = data["id"]

    def test_get_all_drivers(self):
        response = requests.get(f"{BASE_URL}/drivers")
        self.assertEqual(response.status_code, 200)
        data = response.json()
        self.assertIsInstance(data, list)

    def test_get_driver_by_id(self):
        self.test_create_driver()  # First, create a driver
        response = requests.get(f"{BASE_URL}/drivers/{self.driver_id}")
        self.assertEqual(response.status_code, 200)
        data = response.json()
        self.assertEqual(data["name"], "John Doe")
        self.assertEqual(data["license_id"], "12345")

    def test_update_driver(self):
        self.test_create_driver()  # First, create a driver
        payload = {
            "id": self.driver_id,
            "name": "John Doe Updated",
            "license_id": "67890"
        }
        response = requests.put(f"{BASE_URL}/drivers/{self.driver_id}", json=payload)
        self.assertEqual(response.status_code, 200)
        data = response.json()
        self.assertEqual(data["name"], "John Doe Updated")
        self.assertEqual(data["license_id"], "67890")

if __name__ == "__main__":
    unittest.main()
