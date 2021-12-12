


var isPalindrome = function(x) {
    if (x < 0 || (x % 10 === 0 && x !== 0)) {
        return false;
    }

let index = 0;
let revNumber = 0;
let j = x;
    


while(j > 0) {
    revNumber = (revNumber * 10) + (j % 10);
    j = ~~(j/10);
    
    console.log('Index:', index);
    index++;
    console.log(`revNumber: ${revNumber} && j: ${j}`);
}

return x === revNumber;
};


var romanNumeralConversion = function (x) {
    let obj = {
        "one": 'I',
        "five": "V",
        "ten": 'X',
        "fifty": 'L',
        'oneHundred': 'C'
    }
    let result = ''
    if(x.toString().length == 1) {
        let n = x - 1;
        if(n != (5-1)) {
            result += 'IV';
            if(n === (10-1)) {
                result += 'IX'
            }
        }else  {
            result = obj.one + obj.five;
            return result;
        }
    }
}