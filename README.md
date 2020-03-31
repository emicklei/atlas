# asuite - command line tool to use the Atlassian Admin SDK

[![Build Status](https://travis-ci.org/emicklei/gsuite.png)](https://travis-ci.org/emicklei/asuite)

## features

- asuite user list

## configuration

Create an asuite.yaml file (optionally in your HOME directory) such as:

    # get this from admin.atlassian.com URL 
    organisationID: 06fb7edf-53e5-4773-8212-f06c6749c145

    # get this from creating an API Key on admin.atlassian.com
    api-key: **********

## run

    asuite user list
    asuite -config my-asuite.yaml user list

&copy; 2020, ernestmicklei.com. MIT License. Contributions welcome.