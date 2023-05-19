# import logging
# import threading
# import time

# def thread_function(name):
#     logging.info("Thread %s: starting", name)
#     time.sleep(2)
#     logging.info("Thread %s: finishing", name)

# if __name__ == "__main__":
#     format = "%(asctime)s: %(message)s"
#     logging.basicConfig(format=format, level=logging.INFO,
#                         datefmt="%H:%M:%S")

#     logging.info("Main    : before creating thread")
#     x = threading.Thread(target=thread_function, args=(1,))
#     y = threading.Thread(target=thread_function, args=(1,))
#     logging.info("Main    : before running thread")
#     x.start()
#     logging.info("Main    : wait for the thread to finish")
#     # x.join()
#     logging.info("Main    : all done")

import threading
from requests import session
from typing import Dict


def do_request(url: str) -> None:
    client_session = session()
    client_request = client_session.get(url)
    print(client_request.text)
    return


if __name__ == "__main__":
    client_1 = threading.Thread(target=do_request, args=("http://localhost:8080/test1",)).start()
    client_2 = threading.Thread(target=do_request, args=("http://localhost:8080/test2",)).start()
