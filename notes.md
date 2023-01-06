# Notes

## Drag and Drop

Wails and WebView don't supply full path when you drag and drop files.
You can make drag and drop work with the following code, but not getting
the full path is a problem.

```javascript
    function attachDragAndDropEvents() {
        $(window).on('drop', function(e) {
            e.preventDefault();
            e.stopPropagation();
            console.log(e)
            // When drag event is attached to document, use
            // e.dataTransfer.files instead of what's below.
            for (let f of e.originalEvent.dataTransfer.files) {
                console.log(f);
            }
            $(e.currentTarget).removeClass('drop-zone-over');
            return false;
        });
        $(window).on('dragover', function(e) {
            e.preventDefault();
            e.stopPropagation();
            $(e.currentTarget).addClass('drop-zone-over');
            return false;
        });
        $(window).on('dragleave', function(e) {
            e.preventDefault();
            e.stopPropagation();
            $(e.currentTarget).removeClass('drop-zone-over');
            return false;
        });
        $(window).on('dragend', function(e) {
            e.preventDefault();
            e.stopPropagation();
            $(e.currentTarget).removeClass('drop-zone-over');
            return false;
        });
        console.log("Attached drag and drop events")
    }
    attachDragAndDropEvents();
```

## File Browser

If we use a file browser, it must be cross-platform and must support 
multiple selections. [Sqweek Dialog](https://github.com/sqweek/dialog)
is cross-platform, but it supports single selection only. It uses
CGO, which we want to avoid if possible.

[Native File Dialog](https://github.com/mlabbe/nativefiledialog) is 
written in C and would need to be wrapped. It also has some open
issues that may prevent us from being able to use it.

[Zenity](https://github.com/ncruces/zenity) looks good, but it has
external system dependencies on Mac and Linux. Awaiting response from
developer re: external dependencies.

The [Fyne](https://github.com/fyne-io/fyne) file dialog does not seem
to support multiple selection. Single clicking opens a folder instead
of selecting it. It's also ugly.

[wxGo](https://github.com/dontpanic92/wxGo) looks like it's a huge
library to wrap.

[Go QT](https://github.com/therecipe/qt) looks to be quite heavy and
daunting.

## In-app File Tree

The best option at this point may be to display an in-app file tree.
This could appear in the left pane, with a drop target in the right pane.
Some users may be confused, however, thinking they could drop files from
the desktop or file browser there.

We also don't know yet if external file systems will appear correctly.
These include shared network drives and services like iCloud and Dropbox
that appear in the normal finder as mounted drives.

Finally, the appearance should be familiar to Windows/Mac/Linux users.
That is, instead of listing the file system from the root, list familiar
folders like Home, Desktop, Documents, etc. Also list mounted/external
drives below the local list, as on Windows/Mac/Linux.

Since we won't need to walk directories, we can use 
[os.ReadDir](https://pkg.go.dev/os#ReadDir). 