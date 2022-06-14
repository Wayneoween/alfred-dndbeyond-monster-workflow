D&DBeyond Workflow for [Alfred 3 and 4](http://www.alfredapp.com)
==============================

You can search monsters on D&DBeyond (`ddb` in alfred).

You don't need to login but you might end up on a monster page you can't see
because you need to have access to the source book its in.

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
translate` option. A notification will show you the current toggle position. All
future searches will have the german monster name in front and the english
monster name next to it.

Updating
--------

It shouldn't be necessary to manually update, but the framework I used to build
this workflow has a meta keyword you can use.

When you type `ddb workflow:update` it should take a moment while it checks the
[github releases page](https://github.com/Wayneoween/alfred-dndbeyond-monster-workflow/releases))
if there is a newer version than what is installed. If yes alfred opens a
window and asks if you want to update.

Troubleshooting
---------------

If you have a list of empty entries in Alfred while searching for a monster the
cache may be not compatible with the new workflow version. It will probably
help if you do a `ddb workflow:delcache` and try again.

Development
-----------

To get a current list of all source books of monsters you can use `curl` and `jq` like so:

```bash
curl "https://www.dnddeutsch.de/tools/json.php?apiv=0.7&o=monster&q=" |jq -r '.monster[].src[]' |sort |uniq
```

<details><summary>Currently available source</summary>

```text
AI
AiME-BRF
AiME-Eria
AiME-RIV
AiME-RRF
AiME-SLH
AiME-WdD
AiME-Wild
AVENT-M
AVENT-W
BGDiA
CC
CM
CoS
CotN
CTH-GHOUL
CTHULHU
D3
DoIP
EBERRON
EGtW
FToD
GGtR
GoS
HotDQ
IDRotF
LMoP
MARGREVE
MC1
MM
MMM
MOoT
MTGAFR
MToF
Myth-AdDM
Myth-Held
Myth-Saga
OotA
PotA
RAGNAROK
RoT
SCC
SKT
SRD
STRANGE
TalDorei
TDR
ToA
ToB
ToB2
TYP
VGM
VRGtR
WbtW
WDH
WDMM
```

Of those we have to ignore the following:

```text
AiME-BRF
AiME-Eria
AiME-RIV
AiME-RRF
AiME-SLH
AiME-WdD
AiME-Wild
AVENT-M
AVENT-W
CC
CTH-GHOUL
CTHULHU
D3
MARGREVE
MTGAFR
Myth-AdDM
Myth-Held
Myth-Saga
RAGNAROK
STRANGE
ToB
ToB2
```

So we get only these:

```text
AI
BGDiA
CM
CoS
CotN
DoIP
EBERRON
EGtW
FToD
GGtR
GoS
HotDQ
IDRotF
LMoP
MC1
MM
MMM
MOoT
MToF
OotA
PotA
RoT
SCC
SKT
SRD
TalDorei
TDR
ToA
TYP
VGM
VRGtR
WbtW
WDH
WDMM
```

</details>

You can use `source env.sh` to set up an environment where you can run the binary also on non `mac os` operating systems and create a temporary working dir in `./testenv`. From there you can query the API for every monster that matches with the letter `a` which should be almost everyone to build a cache file with about 2000 monsters in json.

```bash
./alfred-dndbeyond-monster-workflow A
```

You can then use the cache to build a list of URLs to test if the renaming works or the naming schema has inconsistencies that need to be catched somehow:

```bash
for i in `cat testenv/cache/A_monster_cache.json |jq -r .[].name_en | tr 'A-Z' 'a-z' | tr ' ' '-'`; do echo https://www.dndbeyond.com/monsters/$i ;done
```

To develop on osx you can use the shell script I copied from [here](https://github.com/lilyball/alfred-install-workflow) like this:

```bash
bash install-workflow.sh -v alfred-dndbeyond-monster-workflow icons
```
