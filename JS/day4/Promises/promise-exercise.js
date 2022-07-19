/*
states-
pending-when declared
fulfilled-when resolved(), use .then 
rejected-when rejected(), use .catch
*/

//2
/*
function delayedLowerCase(word){
       return new Promise((resolve,reject)=>{
        setTimeout(()=>resolve(word.toLowerCase()),5000);
       })
}
delayedLowerCase('BEAUTIFUL CODE').then((lowerWord)=>{
    console.log(lowerWord);
})*/

//3 promise chaining- calling .then on previous .then directly

function delayedLowerCase(word){
    return new Promise((resolve,reject)=>{
        if (typeof word!=='string')
            return reject('No Proper Input, rejected')
        setTimeout(()=>{
            resolve(word.toLowerCase());
        },5000);
    })
}
delayedLowerCase(1).then((lowerWord)=>{
 return lowerWord;
}).then((lowerWord)=>{
    console.log(lowerWord+" - "+lowerWord.length);
}).catch(e=>console.log(e))
