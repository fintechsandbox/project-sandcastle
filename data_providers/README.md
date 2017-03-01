Data Provider Code Base
=======================

Each data provider in the sandbox contains a folder with three sub-folders.
All of the solutions here are language agnostic and Members are encouraged to
contribute in which ever language they'd prefer.

* Each Data provider folder has three sub-folders:
  * **API​**: Documentation that references product descriptions in the Wiki
  * **Issues​**: Solutions to known bugs and references to solutions on StackOverflow
  * **Integrations​**: Solutions to convert external data sources into useful formats for databases, visualizations, etc...

## Adding a new Data provider ##
Copy the contents of the `./data_providers/__PROVIDER_TEMPLATE__` to the new
folder replacing `__PROVIDER_TEMPLATE__` with the appropriate name of the data provider (see naming conventions below)

### Data Provider Naming Conventions ###
 * lowercase
 * replace spaces and periods with underscores:  ` ` -> `_`
 * remove commas and repeating underscores but keep other punctuation

Examples:
 * `SIX Financial Information` -> `six_financial_information`
 * `Morningstar, Inc.` -> `morningstar_inc`

 Example Code implementing the above naming convention
```python
def naming_convention(full_name):
  import re
  N = len(full_name)
  full_name = full_name.lower().strip().replace(' ','_', N).replace(',','', N).replace('.','_', N)
  full_name = re.sub('__*', '_', full_name) # remove duplicate underscores
  full_name = re.sub('_$', '', full_name) # remove leading underscore
  full_name = re.sub('^_', '', full_name) # remove trailing underscore
  return full_name
```
