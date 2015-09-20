# psGoWebApp

Pluralsight course.

http://www.pluralsight.com/courses/creating-web-applications-go

Reminding myself how to use Git and Github at the same time.

## Known Issues during Module 5: MVC The Controller Layer - Part 1

 1. All data is hard-coded.
 2. Only one category (Juices) has been implemented.
 3. GZip compression is active but both Chrome and Firefox claim it's not working.

## Observations about Go

The name of the source file has no bearing on how the contents is referred to in other source files. Any public methods will be visible to a consumer once the package is imported.

I need to read up on function casting to understand difference between func (this \*struct) and func (this type).
