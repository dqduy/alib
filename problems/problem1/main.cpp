using namespace std;

#include<iostream>
#include "../../lib/ArrayList.h"
#include "../../lib/util/ListUtil.h"

int main()
{
    ArrayList a;
    ListUtil::readList(&a);

    a.add(562);
    a.add(-2);

    ListUtil::printList(&a);
    cout << "Size of list: " << a.size() << endl;
    //printf("hello\n");//<< "hello";       	
    return 0;
}
