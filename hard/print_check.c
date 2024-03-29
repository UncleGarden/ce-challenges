#include <stdbool.h>
#include <stdio.h>
#include <string.h>

char *str0[] = {
	"", "One", "Two", "Three", "Four", "Five",
	"Six", "Seven", "Eight", "Nine"
};

char *str1[] = {
	"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen",
	"Sixteen", "Seventeen", "Eighteen", "Nineteen"
};

char *str2[] = {
	"Twenty", "Thirty", "Forty", "Fifty",
	"Sixty", "Seventy", "Eighty", "Ninety"
};

char *str3[] = {
	"Hundred", "Thousand", "Million"
};

bool wrd(int a1, int a2, int a3) {
	if (a1 + a2 + a3 == 0)
		return false;
	if (a1)
		printf("%s%s", str0[a1], str3[0]);
	if (a2 > 1)
		printf("%s%s", str2[a2-2], str0[a3]);
	else if (a2)
		printf("%s", str1[a3]);
	else
		printf("%s", str0[a3]);
	return true;
}

int main(int argc, char *argv[])
{
	FILE *fp;
	int b;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &b) != EOF) {
		int c[9] = {0};
		int i = 8;
		if (b == 0) {
			printf("Zero");
		} else {
			while (b && i >= 0) {
				c[i] = b % 10;
				b /= 10;
				i--;
			}
			if (wrd(c[0], c[1], c[2]))
				printf("%s", str3[2]);
			if (wrd(c[3], c[4], c[5]))
				printf("%s", str3[1]);
			wrd(c[6], c[7], c[8]);
		}
		printf("Dollars\n");
	}
	return 0;
}
