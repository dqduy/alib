#include "../ArrayList.h"
#include "ListUtil.h"
#include<iostream>

using namespace std;

ArrayList* ListUtil::readList(ArrayList* list)
{
    for(int index = 0; index < 4; ++index)
        list->add(index);
    return list;
}

ArrayList* ListUtil::readListWithIncrease(ArrayList* list)
{
    for(int index = 0; index < 21; ++index)
        list->add(index);
    return list;
}

void ListUtil::printList (ArrayList* list)
{
    cout<< "List---" << endl;

    for(int index = 0; index < list->size(); ++index)
    {
        cout<< list->get(index) << " ";
    }

    cout<< endl;
}
