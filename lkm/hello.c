#include<linux/init.h>
#include<linux/module.h>
MODULE_LICENSE("GPL");

static void callUsermodeExecutable(void) {
    int r;
    char *argv[] = { PATH_TO_MAIN, NULL };
    char *envp[] = { NULL };
    r = call_usermodehelper(argv[0], argv, envp, UMH_WAIT_PROC);
    printk("call_usermodehelper returned: %d\n", r);
}

int init_module(void) {
    printk(KERN_ALERT "Hello world!\n");
    callUsermodeExecutable();
    return 0;
}

void cleanup_module(void){
    printk(KERN_ALERT "Goodbye cruel world\n");
}
