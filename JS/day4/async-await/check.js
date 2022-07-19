async function f(){
    try{
        let a=await new Promise((resolve,reject)=>{
                            if (true)
                                return reject("testing out")
                            })
    }
    catch(e){
        console.log(e)
    }
            
}
f()