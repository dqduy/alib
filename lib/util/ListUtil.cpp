#include "../ArrayList.h"
#include "ListUtil.h"
#include<iostream>

using namespace std;

ArrayList* ListUtil::readList(ArrayList* list)
{
	list->add(3);
	list->add(5);
	list->add(15);
	list->add(7);
	list->add(73);
	return list;
}

void ListUtil::printList (ArrayList* list)
{
	cout<<"List---"<<endl;
	
	for(int index = 0; index < list->size(); ++index)
	{
		cout<< list->get(index) << " ";
	}

        cout << endl;
}
