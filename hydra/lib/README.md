# Laniakea Python Library

The laniakea python lib is the core of the hydra compute component.
It includes 
* A CLI to manage auth tokens and register flows
* A task library
* A huge amount of data libraries for "batteries included" data pipelines

## Available Libraries (Reader's digest)

### IO Libraries

#### Http APIs, downloads and HTML scraping
* [requests](https://requests.readthedocs.io/en/master/) Python HTTP for Humans.
* [lxml](https://lxml.de/) Powerful and Pythonic XML processing library combining libxml2/libxslt with the ElementTree API.
* [beautifulsoup](https://www.crummy.com/software/BeautifulSoup/bs4/doc/) Screen-scraping library

#### Database drivers
* [cx-Oracle](https://oracle.github.io/python-cx_Oracle/) Python interface to Oracle
* [ctds](https://zillow.github.io/ctds/) DB API 2.0-compliant driver for SQL Server
* [psycopg2](https://www.psycopg.org/docs/) Python-PostgreSQL Database Adapter
* [mysqlclient](https://mysqlclient.readthedocs.io/) Python interface to MySQL
* [azure-cosmos](https://azuresdkdocs.blob.core.windows.net/$web/python/azure-cosmos/4.1.0/index.html) Microsoft Azure Cosmos Client Library for Python

#### Messaging middleware drivers
* [kafka-python](https://kafka-python.readthedocs.io/en/master/index.html) Pure Python client for Apache Kafka
* [azure-servicebus](https://docs.microsoft.com/en-us/python/api/azure-servicebus/azure.servicebus?view=azure-python) Microsoft Azure Service Bus Client Library for Python
* [google-cloud-pubsub](https://googleapis.dev/python/pubsub/latest/index.html) Google Cloud Pub/Sub API client library

### Data manipulation and analytics
* [pandas](https://pandas.pydata.org/docs/user_guide/index.html) Powerful data structures for data analysis, time series, and statistics
* [vaex](https://docs.vaex.io/en/latest/index.html) Out-of-Core DataFrames to visualize and explore big tabular datasets
* [tensorflow](https://www.tensorflow.org/api_docs/python/tf) TensorFlow is an open source machine learning framework for everyone. 
* [xgboost](https://xgboost.readthedocs.io/en/stable/) XGBoost Python Package
* [scikit-learn](https://scikit-learn.org/stable/install.html) A set of python modules for machine learning and data mining

### Reporting libraries
* [dash](https://dash.plotly.com/) A Python framework for building reactive web-apps.
* [XlsxWriter](https://xlsxwriter.readthedocs.io/) A Python module for creating Excel XLSX files.
* [sendgrid](https://github.com/sendgrid/sendgrid-python) Twilio SendGrid library for Python
* [pyppeteer](https://pyppeteer.github.io/pyppeteer/) Headless chrome/chromium automation library (unofficial port of puppeteer)

## Available Libraries (Full versioned list)

```
absl-py                       0.10.0      Abseil Python Common Libraries, see https://github.com/abseil/abseil-py.
adal                          1.2.4       Note: This library is already replaced by MSAL Python, available here: https://pypi.org/project/msal/ .ADAL Python remains available here as a legacy. The ADAL for Python library mak...
aiohttp                       3.6.2       Async http client/server framework (asyncio)
aplus                         0.11.0      UNKNOWN
appdirs                       1.4.4       A small Python module for determining appropriate platform-specific dirs, e.g. a "user data dir".
argon2-cffi                   20.1.0      The secure Argon2 password hashing algorithm.
astropy                       4.0.1.post1 Community-developed python astronomy tools
astunparse                    1.6.3       An AST unparser for Python
async-timeout                 3.0.1       Timeout context manager for asyncio programs
attrs                         20.1.0      Classes Without Boilerplate
azure-common                  1.1.25      Microsoft Azure Client Library for Python (Common)
azure-core                    1.8.0       Microsoft Azure Core Library for Python
azure-cosmos                  4.1.0       Microsoft Azure Cosmos Client Library for Python
azure-servicebus              0.50.3      Microsoft Azure Service Bus Client Library for Python
backcall                      0.2.0       Specifications for callback functions passed in to an API
bcrypt                        3.2.0       Modern password hashing for your software and your servers
beautifulsoup4                4.9.1       Screen-scraping library
bleach                        3.1.5       An easy safelist-based HTML-sanitizing tool.
bokeh                         2.1.1       Interactive plots and applications in the browser from Python
boto3                         1.14.47     The AWS SDK for Python
botocore                      1.17.47     Low-level, data-driven core of boto 3.
bqplot                        0.12.16     Interactive plotting for the Jupyter notebook, using d3.js and ipywidgets.
branca                        0.4.1       Generate complex HTML+JS pages with Python
brotli                        1.0.7       Python bindings for the Brotli compression library
cachetools                    4.1.1       Extensible memoizing collections and decorators
certifi                       2020.6.20   Python package for providing Mozilla's CA Bundle.
cffi                          1.14.2      Foreign Function Interface for Python calling C code.
chardet                       3.0.4       Universal encoding detector for Python 2 and 3
click                         7.1.2       Composable command line interface toolkit
cloudpickle                   1.5.0       Extended pickling support for Python objects
croniter                      0.3.34      croniter provides iteration for datetime object with cron like format
cryptography                  3.0         cryptography is a package which provides cryptographic recipes and primitives to Python developers.
ctds                          1.12.0      DB API 2.0-compliant driver for SQL Server
cx-oracle                     8.0.0       Python interface to Oracle
cycler                        0.10.0      Composable style cycles
dash                          1.14.0      A Python framework for building reactive web-apps. Developed by Plotly.
dash-core-components          1.10.2      Core component suite for Dash
dash-html-components          1.0.3       Vanilla HTML components for Dash
dash-renderer                 1.6.0       Front-end component renderer for Dash
dash-table                    4.9.0       Dash table
dask                          2.23.0      Parallel PyData with Task Scheduling
decorator                     4.4.2       Decorators for Humans
defusedxml                    0.6.0       XML bomb protection for Python stdlib modules
distributed                   2.23.0      Distributed scheduler for Dask
docker                        4.3.1       A Python library for the Docker Engine API.
docutils                      0.15.2      Docutils -- Python Documentation Utilities
entrypoints                   0.3         Discover and load entry points from installed packages.
fastavro                      0.24.2      Fast read/write of AVRO files
flask                         1.1.2       A simple framework for building complex web applications.
flask-compress                1.5.0       Compress responses in your Flask app with gzip or brotli.
fsspec                        0.8.0       File-system specification
future                        0.18.2      Clean single-source support for Python 3 and 2
gast                          0.3.3       Python AST that abstracts the underlying Python version
gcsfs                         0.7.0       Convenient Filesystem interface over GCS
gitdb                         4.0.5       Git Object Database
gitpython                     3.1.7       Python Git Library
google-api-core               1.22.1      Google API client core library
google-auth                   1.20.1      Google Authentication Library
google-auth-oauthlib          0.4.1       Google Authentication Library
google-cloud-bigquery         1.27.2      Google BigQuery API client library
google-cloud-bigquery-storage 1.0.0       BigQuery Storage API API client library
google-cloud-core             1.4.1       Google Cloud API client core library
google-cloud-datacatalog      1.0.0       Google Cloud Data Catalog API API client library
google-cloud-pubsub           1.7.0       Google Cloud Pub/Sub API client library
google-cloud-storage          1.30.0      Google Cloud Storage API client library
google-crc32c                 0.1.0       A python wrapper of the C library 'Google CRC32C'
google-pasta                  0.2.0       pasta is an AST-based Python refactoring library
google-resumable-media        0.7.1       Utilities for Google Media Downloads and Resumable Uploads
googleapis-common-protos      1.52.0      Common protobufs used in Google APIs
grpc-google-iam-v1            0.12.3      GRPC library for the google-iam-v1 service
grpcio                        1.31.0      HTTP/2-based RPC framework
h5py                          2.10.0      Read and write HDF5 files from Python
heapdict                      1.0.1       a heap with decrease-key and increase-key operations
idna                          2.10        Internationalized Domain Names in Applications (IDNA)
ipydatawidgets                4.0.1       A set of widgets to help facilitate reuse of large datasets across widgets
ipykernel                     5.3.4       IPython Kernel for Jupyter
ipyleaflet                    0.13.3      A Jupyter widget for dynamic Leaflet maps
ipympl                        0.5.7       Matplotlib Jupyter Extension
ipython                       7.17.0      IPython: Productive Interactive Computing
ipython-genutils              0.2.0       Vestigial utilities from IPython
ipyvolume                     0.5.2       IPython widget for rendering 3d volumes
ipyvue                        1.4.0       Jupyter widgets base for Vue libraries
ipyvuetify                    1.5.1       Jupyter widgets based on vuetify UI components
ipywebrtc                     0.5.0       WebRTC for Jupyter notebook/lab
ipywidgets                    7.5.1       IPython HTML widgets for Jupyter
isodate                       0.6.0       An ISO 8601 date/time/duration parser and formatter
itsdangerous                  1.1.0       Various helpers to pass data to untrusted environments and back.
jedi                          0.17.2      An autocompletion tool for Python that can be used for text editors.
jinja2                        2.11.2      A very fast and expressive template engine.
jmespath                      0.10.0      JSON Matching Expressions
joblib                        0.16.0      Lightweight pipelining: using Python functions as pipeline jobs.
json5                         0.9.5       A Python implementation of the JSON5 data format.
jsonschema                    3.2.0       An implementation of JSON Schema validation for Python
jupyter-client                6.1.6       Jupyter protocol implementation and client libraries
jupyter-core                  4.6.3       Jupyter core package. A base package on which Jupyter projects rely.
jupyterlab                    2.2.5       The JupyterLab notebook server extension.
jupyterlab-server             1.2.0       JupyterLab Server
kafka-python                  2.0.1       Pure Python client for Apache Kafka
keras-preprocessing           1.1.2       Easy data preprocessing and data augmentation for deep learning models
kiwisolver                    1.2.0       A fast implementation of the Cassowary constraint solver
llvmlite                      0.34.0      lightweight wrapper around basic LLVM functionality
lxml                          4.5.2       Powerful and Pythonic XML processing library combining libxml2/libxslt with the ElementTree API.
markdown                      3.2.2       Python implementation of Markdown.
markupsafe                    1.1.1       Safely add untrusted strings to HTML/XML markup.
marshmallow                   3.7.1       A lightweight library for converting complex datatypes to and from native Python datatypes.
marshmallow-oneofschema       2.0.1       marshmallow multiplexing schema
matplotlib                    3.3.1       Python plotting package
mistune                       0.8.4       The fastest markdown parser in pure Python
msgpack                       1.0.0       MessagePack (de)serializer.
msrest                        0.6.18      AutoRest swagger generator Python client runtime.
msrestazure                   0.6.4       AutoRest swagger generator Python client runtime. Azure-specific module.
multidict                     4.7.6       multidict implementation
mypy-extensions               0.4.3       Experimental type system extensions for programs checked with the mypy typechecker.
mysqlclient                   2.0.1       Python interface to MySQL
natsort                       7.0.1       Simple yet flexible natural sorting in Python.
nbconvert                     5.6.1       Converting Jupyter Notebooks
nbformat                      5.0.7       The Jupyter Notebook format
nest-asyncio                  1.4.0       Patch asyncio to allow nested event loops
notebook                      6.1.3       A web-based notebook environment for interactive computing
numba                         0.51.0      compiling Python code using LLVM
numpy                         1.18.5      NumPy is the fundamental package for array computing with Python.
oauthlib                      3.1.0       A generic, spec-compliant, thorough implementation of the OAuth request-signing logic
opt-einsum                    3.3.0       Optimizing numpys einsum function
packaging                     20.4        Core utilities for Python packages
pandas                        1.1.1       Powerful data structures for data analysis, time series, and statistics
pandocfilters                 1.4.2       Utilities for writing pandoc filters in python
paramiko                      2.7.1       SSH2 protocol library
parso                         0.7.1       A Python Parser
pendulum                      2.1.2       Python datetimes made easy
pexpect                       4.8.0       Pexpect allows easy control of interactive console applications.
pickleshare                   0.7.5       Tiny 'shelve'-like database with concurrency support
pillow                        7.2.0       Python Imaging Library (Fork)
plotly                        4.9.0       An open-source, interactive data visualization library for Python
prefect                       0.13.1      The Prefect Core automation and scheduling engine.
progressbar2                  3.51.4      A Python Progressbar library to provide visual (yet text based) progress to long running operations.
prometheus-client             0.8.0       Python client for the Prometheus monitoring system.
prompt-toolkit                3.0.6       Library for building powerful interactive command lines in Python
protobuf                      3.13.0      Protocol Buffers
psutil                        5.7.2       Cross-platform lib for process and system monitoring in Python.
psycopg2                      2.8.5       psycopg2 - Python-PostgreSQL Database Adapter
ptyprocess                    0.6.0       Run a subprocess in a pseudo terminal
pyarrow                       1.0.1       Python library for Apache Arrow
pyasn1                        0.4.8       ASN.1 types and codecs
pyasn1-modules                0.2.8       A collection of ASN.1-based protocols modules.
pycparser                     2.20        C parser in Python
pyee                          7.0.2       A port of node.js's EventEmitter to python.
pygments                      2.6.1       Pygments is a syntax highlighting package written in Python.
pyjwt                         1.7.1       JSON Web Token implementation in Python
pynacl                        1.4.0       Python binding to the Networking and Cryptography (NaCl) library
pyopenssl                     19.1.0      Python wrapper module around the OpenSSL library
pyparsing                     2.4.7       Python parsing module
pyppeteer                     0.2.2       Headless chrome/chromium automation library (unofficial port of puppeteer)
pyrsistent                    0.16.0      Persistent/Functional/Immutable data structures
pysftp                        0.2.9       A friendly face on SFTP
python-box                    4.2.3       Advanced Python dictionaries with dot notation access
python-dateutil               2.8.1       Extensions to the standard Python datetime module
python-http-client            3.3.0       HTTP REST client, simplified for Python
python-slugify                4.0.1       A Python Slugify application that handles Unicode
python-utils                  2.4.0       Python Utils is a module with some convenient utilities not included with the standard Python install
pythreejs                     2.2.0       Interactive 3d graphics for the Jupyter notebook, using Three.js from Jupyter interactive widgets.
pytz                          2020.1      World timezone definitions, modern and historical
pytzdata                      2020.1      The Olson timezone database for Python.
pyyaml                        5.3.1       YAML parser and emitter for Python
pyzmq                         19.0.2      Python bindings for 0MQ
requests                      2.24.0      Python HTTP for Humans.
requests-oauthlib             1.3.0       OAuthlib authentication support for Requests.
retrying                      1.3.3       Retrying
rsa                           4.6         Pure-Python RSA implementation
ruamel.yaml                   0.16.10     ruamel.yaml is a YAML parser/emitter that supports roundtrip preservation of comments, seq/map flow style, and map key order
ruamel.yaml.clib              0.2.0       C version of reader, parser and emitter for ruamel.yaml derived from libyaml
s3fs                          0.2.2       Convenient Filesystem interface over S3
s3transfer                    0.3.3       An Amazon S3 Transfer Manager
scikit-learn                  0.23.2      A set of python modules for machine learning and data mining
scipy                         1.4.1       SciPy: Scientific Library for Python
semver                        2.10.2      Python helper for Semantic Versioning (http://semver.org/)
send2trash                    1.5.0       Send file to trash natively under Mac OS X, Windows and Linux.
sendgrid                      6.4.6       Twilio SendGrid library for Python
six                           1.15.0      Python 2 and 3 compatibility utilities
smmap                         3.0.4       A pure Python implementation of a sliding window memory map manager
sortedcontainers              2.2.2       Sorted Containers -- Sorted List, Sorted Dict, Sorted Set
soupsieve                     1.9.6       A modern CSS selector implementation for Beautiful Soup.
sqlalchemy                    1.3.19      Database Abstraction Library
starkbank-ecdsa               1.0.0       A lightweight and fast pure python ECDSA library
tabulate                      0.8.7       Pretty-print tabular data
tblib                         1.7.0       Traceback serialization library.
tensorboard                   2.3.0       TensorBoard lets you watch Tensors Flow
tensorboard-plugin-wit        1.7.0       What-If Tool TensorBoard plugin.
tensorflow                    2.3.0       TensorFlow is an open source machine learning framework for everyone.
tensorflow-estimator          2.3.0       TensorFlow Estimator.
termcolor                     1.1.0       ANSII Color formatting for output in terminal.
terminado                     0.8.3       Terminals served to xterm.js using Tornado websockets
testpath                      0.4.4       Test utilities for code working with files and commands
text-unidecode                1.3         The most basic Text::Unidecode port
threadpoolctl                 2.1.0       threadpoolctl
toml                          0.10.1      Python Library for Tom's Obvious, Minimal Language
toolz                         0.10.0      List processing tools and functional utilities
tornado                       6.0.4       Tornado is a Python web framework and asynchronous networking library, originally developed at FriendFeed.
tqdm                          4.48.2      Fast, Extensible Progress Meter
traitlets                     4.3.3       Traitlets Python config system
traittypes                    0.2.1       Scipy trait types
typing-extensions             3.7.4.2     Backported and Experimental Type Hints for Python 3.5+
uamqp                         1.2.10      AMQP 1.0 Client Library for Python
urllib3                       1.25.10     HTTP library with thread-safe connection pooling, file post, and more.
vaex                          3.0.0       Out-of-Core DataFrames to visualize and explore big tabular datasets
vaex-arrow                    0.5.1       Arrow support for vaex
vaex-astro                    0.7.0       Astronomy related transformations and FITS file support
vaex-core                     2.0.3       Core of vaex
vaex-hdf5                     0.6.0       hdf5 file support for vaex
vaex-jupyter                  0.5.2       Jupyter notebook and Jupyter lab support for vaex
vaex-ml                       0.9.0       Machine learning support for vaex
vaex-server                   0.3.1       Webserver and client for vaex for a remote dataset
vaex-viz                      0.4.0       Visualization for vaex
wcwidth                       0.2.5       Measures the displayed width of unicode strings in a terminal
webencodings                  0.5.1       Character encoding aliases for legacy web content
websocket-client              0.57.0      WebSocket client for Python. hybi13 is supported.
websockets                    8.1         An implementation of the WebSocket Protocol (RFC 6455 & 7692)
werkzeug                      1.0.1       The comprehensive WSGI web application library.
wheel                         0.35.1      A built-package format for Python
widgetsnbextension            3.5.1       IPython HTML widgets for Jupyter
wrapt                         1.12.1      Module for decorators, wrappers and monkey patching.
xarray                        0.16.0      N-D labeled arrays and datasets in Python
xgboost                       1.1.1       XGBoost Python Package
xlsxwriter                    1.3.3       A Python module for creating Excel XLSX files.
yarl                          1.5.1       Yet another URL library
zict                          2.0.0       Mutable mapping tools
```