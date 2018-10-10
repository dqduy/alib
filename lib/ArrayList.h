#ifndef _ArrayList_H
#define _ArrayList_H

#include<cstdint>

class ArrayList
{
    private:
        int64_t capacity;
        int64_t totalItems;
        int     blockSize;      //Space added to old space
        int*    list;           //Hold head of list
    
        void increaseSpace(int space);
        void decreaseSpace(int space);
    public: 
        ArrayList();
        ArrayList(int64_t capacity);
        ArrayList(ArrayList& obj);
        bool isEmpty();
        bool isFull();                          
        void clear();                                   //Reset current list
        void insert(int64_t location, int item);        //Add an item to the list in location (zero-based)
        void add(ArrayList& list);                      //For future
        void add(int item);                             //Add an item to the list
        void remove(int64_t location);                  //Remove an item with specific index
        int64_t indexOf(int item);  
        int get(int64_t location);
        int64_t size();	

};

#endif

