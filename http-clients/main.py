import threading
from requests import session
from typing import Dict


def do_request(url: str, client_id: int) -> None:
    client_session = session()
    client_request = client_session.get(url)
    print(client_id, client_request.text)
    print(client_id, client_request.headers)
    return


if __name__ == "__main__":
    client_1 = threading.Thread(target=do_request, args=("http://localhost:2004/test1", 1)).start()
    client_2 = threading.Thread(target=do_request, args=("http://localhost:2004/test1", 2)).start()
