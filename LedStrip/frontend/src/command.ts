export const HOST = 'http://localhost:8121'

export const sendCommand = (cmd: string) => {
    let formData = new FormData();
    formData.append('command', cmd);

    fetch(HOST + '/exec', {
        body: formData,
        method: 'POST',
    })
}

export const convertColor = (col: string): string => {
    return '0x' + col.substr(1, 2) + col.substr(5, 2) + col.substr(3, 2);
}