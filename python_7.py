# coding=utf-8

class Animal:
    '''this is a help doc'''
    name = 'no name'
    owner = 'no owner'
    noise = 'no noise'

    def __init__(self, name, owner, noise):
        self.name = name
        self.owner = owner
        self.noise = noise

    def __del__(self):
        print 'destroy'

    def printName(self):
        print self.name

    def printOwner(self):
        print self.owner

    def printNoise(self):
        print self.noise

    def __privateFun(self):
        print 'hard to find'


class Dog(Animal):
    def __init__(self, name, owner, noise):
        Animal.__init__(self, name, owner, noise)

    def printNoise(self):
        print 'woof'


class Cat(Animal):
    def __init__(self, name, owner, noise):
        Animal.__init__(self, name, owner, noise)


class Dot(Dog, Cat):
    def __init__(self, name, owner, noise):
        Cat.__init__(self, name, owner, noise)
        Dog.__init__(self, name, owner, noise)


def main():
    aDot = Dot('papi', 'weizheng', 'bark')
    aDot.printNoise()
    print issubclass(Cat, Animal)
    print Dot.__bases__
    print aDot.__dict__


if __name__ == '__main__':
    main()

print Animal.__doc__













