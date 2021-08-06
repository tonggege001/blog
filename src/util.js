export function GetDateFromString(dateStr){
    let dateArr = dateStr.split("-")
    let year = parseInt(dateArr[0])
    if (dateArr[1][0] == '0') {
        dateArr[1] = dateArr[1].substr(1)
    }
    let month = parseInt(dateArr[1])

    if (dateArr[2][0] == '0') {
        dateArr[2] = dateArr[2].substr(1)
    }
    let day = parseInt(dateArr[2])
    return {
        year: year,
        month: month,
        day: day
    }
}
