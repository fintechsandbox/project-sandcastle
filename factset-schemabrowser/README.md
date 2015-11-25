#Browsing Factset Data

## Usage
To run the last compiled version of the schema from Factset's custom schema format, you will need to serve this directory from a web server.
If you have python running on your machine

```
cd project-sandcastle/factset-schemabrowser
python -m SimpleHTTPServer 8000
```

Then navigate to http://localhost:8000/index.html to browse and search the data.

## Updating to the latest version of the data

### Pre-requisites 
*Change directory to project-sandcastle/factset-schemabrowser
*Install node and npm
* Download the latest version of the schema file from Factset FTP feed. The file is located at `sandbox-files/schema/spl_v2_schema.txt`
** The sample file is uploaded in the project-sandcastle/factset-schemabrowser/data directory

### Generating a new version of the schema
Run the following command to generate a new version of the schema
```
node generateSchema.js <full-path to downloaded file>
```

Then navigate to http://localhost:8000/index.html
