#include "ArrayList.h"

void ArrayList::increaseSpace(int space)
{
}

void ArrayList::decreaseSpace(int space) 
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
    this->capacity    = capacity;
    this->totalItems  = 0;
    this->blockSize   = 124;
    list = new int[this->capacity];
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
    list = new int[capacity];
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
        totalItems++;  
    }

    if(totalItems == capacity)
        increaseSpace(blockSize);
}

void ArrayList::remove(int64_t location)
{
    if(location >= 0 && location < totalItems)
    {
        for(int index = location; index <= totalItems - 2; ++index)
            list[index] = list[index + 1];
        totalItems--;
    }

    if(totalItems == 0)
        decreaseSpace(capacity);

} 

int64_t ArrayList::indexOf(int item)
{
    for(int index = 0; index < totalItems; ++index) 
    {
        if(list[index] == item)
            return index;
    }

    return -1;
} 

int ArrayList::get(int64_t location)
{
    return list[location];
} 

int64_t ArrayList::size()
{
    return totalItems;
}
