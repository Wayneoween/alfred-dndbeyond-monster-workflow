D&DBeyond Workflow for [Alfred 3](http://www.alfredapp.com)
==============================

You can search monsters on D&DBeyond (`ddb`).

You don't need to login but you might end up on a monster page you can't see
because you didn't buy the source its in.

**[DOWNLOAD](https://github.com/Wayneoween/alfred-dndbeyond-monster-workflow/releases)**

![Demo GIF](demo.gif)

Usage
-----

Just type `ddb goblin` and you have a list of D&D monsters and their basic properties.

Pressing Enter on a result leads you to its respective complete statblock on
[](https://dndbeyond.com). Should you not own the necessary source book there
you will get redirected to the store page.

If you want the german name of the monster, as well as the page you can find it
in the german source books, you can type `ddb` and select the `Toggle
translate` field. A notification will show you the current toggle position. All
future searches will have the german monster name in front and the english
monster name next to it.

Updating
--------

It shouldn't be necessary to manually update, but the framework I used to build
this workflow has a meta keyword you can use.

When you type `ddb workflow:update` it should take a moment while it checks the
[github releases
page](](https://github.com/Wayneoween/alfred-dndbeyond-monster-workflow/releases))#
if there is a newer version than what is installed. If yes alfred opens a
window and asks if you want to update.

Troubleshooting
---------------

If you have a list of empty entries in Alfred while searching for a monster the
cache may be not compatible with the new workflow version. It will probably
help if you do a `ddb workflow:delcache` and try again.

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

![Workflow Screenshot](screenshot.png)
