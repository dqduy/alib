#include "ArrayList.h"

ArrayList::increaseSpace(int space)
{
}

ArrayList::decreaseSpace(int space) 
{
}

ArrayList::ArrayList() 
{
    capacity    = 20;
    totalItems  = 0;
    blockSize   = 124;
    list = new int[capacity];
}

ArrayList::ArrayList(int64_t capacity) 
{
    this.capacity    = capacity;
    this.totalItems  = 0;
    this.blockSize   = 124;
    list = new int[this.capacity];
}

ArrayList::ArrayList(ArrayList& obj) 
{
}

bool ArrayList::isEmpty()
{
    if(totalItems == 0) return true;
    else return false;
}

bool ArrayList::isFull()
{
    if(totalItems == capacity) return true;
    else return false;
}

void ArrayList::clear()
{
    delete[] list;
    totalItems = 0;
}

void ArrayList::insert(int64_t location, int item)
{
}

void ArrayList::add(ArrayList& obj)
{
}

void ArrayList::add(int item)
{
    if(!isFull())
    {
        list[totalItems] = item;
        //if(totalItems ==)
        totalItems++;
        //Add item 
        
            
    }
}

void ArrayList::remove(int64_t location)
{

} 

int64_t ArrayList::indexOf(int item)
{
    return 0;
} 

int ArrayList::get(int64_t location)
{
    return list[localtion];
} 

int64_t ArrayList::size()
{
    return totalItems;
}
