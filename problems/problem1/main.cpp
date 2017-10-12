using namespace std;

#include<iostream>
#include "../../lib/ArrayList.h"
#include "../../lib/util/ListUtil.h"

int main()
{
    ArrayList a;
    ListUtil::readList(&a);
    ListUtil::printList(&a);
    cout << "Size of list: " << a.size() << endl;
    a.clear();

    cout << "\n";

    ListUtil::readListWithIncrease(&a);
    ListUtil::printList(&a);
    cout << "Size of list: " << a.size() << endl;

    return 0;
}
