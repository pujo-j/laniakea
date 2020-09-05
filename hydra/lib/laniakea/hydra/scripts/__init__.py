import logging
import os


def init_logging():
    if os.getenv("DEBUG"):
        print("Debug mode enabled")
        logging.getLogger("cli").setLevel(logging.DEBUG)
        logging.getLogger("laniakea.hydra.loader").setLevel(logging.DEBUG)
