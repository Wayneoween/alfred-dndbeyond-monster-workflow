D&DBeyond Workflow for [Alfred 3](http://www.alfredapp.com)
==============================

You can search monsters on D&DBeyond (`ddb`).

You don't need to login but you might end up on a monster page you can't see because you didn't buy the source its in.

**[DOWNLOAD](https://github.com/Wayneoween/alfred-dndbeyond-monster-workflow/releases)**

![Workflow Screenshot](screenshot.png)

Development
-----------

This project uses `golang v1.12` and the preliminary version of its [version
management](https://github.com/golang/go/wiki/Modules) called modules. To make
everything work I use the libraries [`awgo`](https://github.com/deanishe/awgo)
to interact with alfred and [`colly`](http://go-colly.org) to scrape the
[](https://dndbeyond.com) search results with the monster filter applied.

Also the very nice
[`workflow-install.py`](https://gist.github.com/deanishe/35faae3e7f89f629a94e)
is from [deanishe](https://github.com/deanishe).

Development workflow is like this:

1. Make changes
2. `python2 workflow-install.py`
3. Test changes in Alfred
4. Repeat
