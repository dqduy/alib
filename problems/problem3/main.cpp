using namespace std;

#include<iostream>
#include "tinyxml2.h"

int main()
{
    tinyxml2::XMLDocument doc;
    doc.LoadFile("sample.html");
    doc.Print();
    return 0;
}
