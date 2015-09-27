# psGoWebApp

Pluralsight course.

http://www.pluralsight.com/courses/creating-web-applications-go

Reminding myself how to use Git and Github at the same time.

## Known Issues during Module 8: Persisting Data

 1. All data is hard-coded.
 2. Only one category (Juices) has been implemented.
 3. Stand_locator doesn't implement the postcode finder. Just returns hardcoded data.
 4. Tests only added for a few types.
 
## Observations from this learning project

The name of the source file has no bearing on how the contents is referred to in other source files. Any public methods will be visible to a consumer once the package is imported.

Need to read up on function casting to understand difference between func (this \*struct) and func (this type).

MVC has proven hard to follow in this project. Naming conventions (matching  concepts have the same name in model/package/controller/converter folders) may be a large part of the issue.
 
