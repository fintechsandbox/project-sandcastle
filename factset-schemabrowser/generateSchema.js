#! /usr/bin/env node

'use strict';

var _ = require('lodash');
var Readline = require('readline');
var Fs = require('fs');

//var userArgs = process.argv.slice(2);
//var companySearch = userArgs[0];


function schemaLineParser(filename, delimiter)
{

	var processor = new ProcessLine();

	var rl = Readline.createInterface({
		input: Fs.createReadStream(filename)
	});

	var rowCtr = 0;
	rl.on('line', function (line) {
		var lineParts = line.split( delimiter || '|' );
		processor.process(lineParts );
		rowCtr+=1;
	});

	rl.on('close', function()
	{
		//console.log(JSON.stringify(processor.tables));
        Fs.writeFile('data/schema.json', JSON.stringify(processor.tables), function(err)
        {
            if(err) return console.log('Something went really wrong... Uggh!');

            console.log('Generated output file in [data/schema.json] ');
            process.exit(0);
        });
	});

};


var Bundle = function(bundlePrefix, directory, feedName)
{
	this.bundlePrefix = bundlePrefix;
	this.directory = directory;
	this.feedName = feedName;
};

var Table = function(tableName, tableDescription, textFilePrefix)
{
	this.tableName = tableName;
	this.bundle = undefined;
	this.tableDescription = tableDescription;
	this.textFilePrefix = textFilePrefix;
	this.primaryKeys  = [];
	this.foreignKeys = [];
	this.columns = [];
};

Table.prototype.linkBundle = function(bundleName, bundleIndex)
{
	this.bundle = bundleIndex[bundleName];
};

var Column = function(colPos, tableName, fieldName, dataType, nullable, desc)
{
	this.colPos = colPos;
	this.tableName = tableName;
	this.fieldName = fieldName;
	this.dataType = dataType;
	this.nullable = nullable;
	this.colDescription = desc;
};

Column.prototype.linkTable = function(tableName, tableIndex)
{
	var table =  tableIndex[tableName];
	table.columns.push(this);
};

var ProcessLine = function(){

	this.bundles = [];
	this.bundleIndex = undefined;
	this.tables = [];
	this.tableIndex = undefined;
};

ProcessLine.prototype.process = function(lineParts){
	var schemaDef = lineParts[0];
	switch(schemaDef)
	{
		case '0' :
			break;
		case '1' :
			// Process Feed definitions
			var bundlePrefix = lineParts[1];
			var directory = lineParts[2];
			var feedName = lineParts[3];
			var bundle = new Bundle(bundlePrefix, directory, feedName);
			this.bundles.push(bundle);
			break;
		case '2':
			//console.log(bundles);
			//Process Table Meta
			var tableName = lineParts[1];
			var bundleName = lineParts[2];
			var textFilePrefix = lineParts[3];
			var tableDesc = lineParts[4];
			var table = new Table(tableName, tableDesc, textFilePrefix);
			if(!this.bundleIndex){
				this.bundleIndex = _.indexBy(this.bundles, 'bundlePrefix');
			}
			table.linkBundle(bundleName, this.bundleIndex);

			this.tables.push(table);
			break;
		case '3':
			//Process Table and Cols
			var colPos = lineParts[1];
			var tableName = lineParts[2];
			var fieldName = lineParts[3];
			var dataType = lineParts[4];
			var nullable = lineParts[5];
			var desc = lineParts[6];
			if(!this.tableIndex){
				this.tableIndex = _.indexBy(this.tables, 'tableName');
			}

			var col = new Column(colPos, tableName, fieldName, dataType,nullable, desc);
			col.linkTable(tableName, this.tableIndex);
			break;
		case '4':
			//Process primary keys
			var tableName = lineParts[1];
			var table = this.tableIndex[tableName];
			table.primaryKeys =  lineParts[2].split(',');
			break;

		case '5':
			//Process foreign keys
			//Child Table Name|Child Field Name(s)|Parent Table Name|Parent Field Names

			break;

	}
};


function existsSync(filename) {
    try {
        Fs.accessSync(filename);
        return true;
    } catch(ex) {
        return false;
    }
}

var argLength = process.argv.length;
if(argLength < 3)
{
	console.log('Usage: node generateSchema.js <full path to schema file>');
	process.exit(1);
}

if(existsSync(process.argv[2]))
{
    schemaLineParser(process.argv[2]);
}
else
{
    console.log('File ['+ process.argv[2] +'] not found. Uggh!');
    console.log('Specify the full path to the Factset schema file');
    process.exit(1);
}


//schemaLineParser('/Users/sanjayvenkat2000/Documents/data/sandbox-files/schema/spl_v2_schema.txt');






