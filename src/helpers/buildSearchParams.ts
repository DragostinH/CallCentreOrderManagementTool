
export function buildSearchParams<T extends object>(data: T): URLSearchParams {
    const params = new URLSearchParams()

    for (const [key, value] of Object.entries(data)) {
        if (value != null && value != '' && value != undefined) {
            params.append(key, String(value))
        }
    }

    return params

}