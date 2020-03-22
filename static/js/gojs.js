function copyClipboard(el) {
    /* Get the text field */
    var copyText = document.getElementById(el)

    /* Select the text field */
    copyText.select();
    copyText.setSelectionRange(0, 99999); /*For mobile devices*/

    /* Copy the text inside the text field */
    document.execCommand("copy");

    /* Alert the copied text */
    M.toast({ html: 'Output copied', displayLength: 1000 })
}

(function ($) {
    $(function () {
        $('.sidenav').sidenav();

        $('.fixed-action-btn').floatingActionButton();

        window.setInterval(function () {
            M.textareaAutoResize($('#output-text'));
        }, 1000);

    }); // end of document ready
})(jQuery); // end of jQuery name space