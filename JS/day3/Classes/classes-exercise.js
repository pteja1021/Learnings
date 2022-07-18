class Queue{
    #queue=[];
    constructor(){
        this.#queue=[]
    }
    enqueue(item){
        this.#queue.push(item);
        console.log(item+" added");
    }
    dequeue(item){
        let a=this.#queue.shift();
        console.log("dequeued item: "+a)
    }
    front(){
        let a=this.#queue[0];
        return a;
    }
    isEmpty(){
        return this.#queue.length===0?true:false;
    }
    toString(){
        console.log(this.#queue)
    }
}
let q= new Queue()
q.enqueue(1)
q.enqueue(2)
console.log(q)

// instanceof is used to check if an object is an instance of a particular class
// usage: q instanceof Queue - returns boolean true