function emailValidator() {
    return function email(value) {
        return (value && !!value.match(/^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/))
    }
}

function requiredValidator() {
    return function required(value) {
        return (value !== undefined && value !== null && value !== '') || 'Enter IP adress'
    }
}

export {
    emailValidator,
    requiredValidator
}
