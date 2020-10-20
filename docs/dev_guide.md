# Laniakea Getting Started

* [Developer Environment](#developer-environment)
* [Using Task Libraries](#using-task-libraries)
    + [Creating a task library](#creating-a-task-library)
    + [Deploying a task library](#deploying-a-task-library)
* [Deploying on a Laniakea norma server](#deploying-on-a-laniakea-norma-server)

## Developer Environment

A prepared container-based dev environment, based on eclipse Theia is provided and can be run with the following docker command

```shell script
docker run -p 3000:3000 pujoj/laniakea-dev:0.2.2
```

Which will expose a web UI based on visual studio code, with a preconfigured hydra dev environment to create and test pipelines on http://localhost:3000/
Inside this environment you can write prefect based pipelines using all [available hydra libraries](https://github.com/pujo-j/laniakea/blob/master/hydra/lib/README.md)

To devlop pipelines, follow the [Prefect core doc](https://docs.prefect.io/core/)

## Using Task Libraries

While you can define tasks in the flow definition python file for simple flows, the single file rule makes it hard to maintain at scale.


Laniakea has a dynamic loading mechanism for task libraries that allows for task versioning.


Any task that can be reused in multiple flows should be extracted in a versioned task library.


Task libraries are detected from the HYDRA_REPO environment variable, which is a comma separated list of repositories.


### Creating a task library

A task library is a collection of python packages, in development you can use a simple directory containing python packages and point to it by adding file://$LIBRARY_PATH_1/ to the HYDRA_REPO list.

### Deploying a task library

Deployment of a task library depends on your repository implementation.


Currently there are two production implementations, both use standard PKZIP compressed files containing python packages.


The zip file MUST include a version number, and the filename MUST be a valid python package name.


Copying it to the Google Cloud Storage repository location, or to a properly configured web base URL is all that is necessary for “publication”, your CI/CD should do it for you !


## Deploying on a Laniakea norma server

The server administrator should have given you a deployment token, address and project name.

Set the following environment variables:

* PREFECT__CLOUD__AUTH_TOKEN : Your deployment auth token 
* LANIAKEA_STORAGE: http://$SERVER_ADDRESS/storage/
* PREFECT__SERVER__ENDPOINT: http://$ SERVER_ADDRESS /gql
* HYDRA_REPO: The task library repository base URL

Where $SERVER_ADDRESS is the Laniakea server address

Once those variables are registered, you just need to use the laniakea command line application to deploy:
```shell script
lnk register -f [YOUR_FLOW_PYTHON_FILE] -p [PROJECT_ID]
```

Where PROJECT_ID is an existing organization project name on the Laniakea server.

