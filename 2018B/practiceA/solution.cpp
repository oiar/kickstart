#include <bits/stdc++.h>
using namespace std;
#define LL long long
LL pn[17] = {1};

LL noNine(LL x) {
  vector<int> num;
  LL t = x, res = 0;
  if(t == 0) {
    num.push_back(0);
  }
  while (t)
  {
    num.push_back(t%10);
    t /= 10;
  }
  int len = num.size();
  for(int i = len - 1; i > 0; i--) {
    res += num[i] * pn[i - 1] * 8;
    if(num[i] == 9) {
      return res;
    }
  }
  x -= num[0];
  for(int i = 0; i <= num[0] && i != 9; i++) {
    if((x + i) % 9) {
      res++;
    }
  }
  return res;
}

int main() {
  int T, o = 0;
  for(int i = 1; i < 18; i++) {
    pn[i] = pn[i - 1]*9;
  }
  cin >> T;
  while (T--)
  {
    LL f, l;
    cin >> f >> l;
    LL res = noNine(l) - noNine(f) + 1;
    printf("Case #%d: %lld\n", ++o, res);
  }
  return 0;
}