chrome.browserAction.onClicked.addListener(() => {
    const input = document.createElement('input');
    input.style.position = 'fixed';
    input.style.opacity = 0;
    document.body.appendChild(input);
    input.focus();
    document.execCommand('paste');
    const url = input.value;

    document.body.removeChild(input);

    if (url.indexOf('youtube.com') !== -1) {
        sendLink(url)

    }
});


function sendLink(url) {
    fetch('http://localhost:9086/link', {
        method: 'POST',
        body: JSON.stringify({url: url}),
        headers: {'Content-Type': 'application/json'}
    }).then(r => console.log('Ok'))
}