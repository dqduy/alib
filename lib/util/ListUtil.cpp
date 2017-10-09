#include "ListUtil.hpp"

using namespace std;

ArrayList* ListUtil::readList(ArrayList* list)
{
	list->add(3);
	list->add(5);
	list->add(15);
	list->add(1, 7);
	list->add(2, 73);
	return list;
}

void ListUtil::printList (ArrayList* list)
{
	cout<<"List---"<<endl;
	
	for(int i = 1; i <= list->size(); i++)
	{
		cout<< list->get(i) << " ";
	}
}
