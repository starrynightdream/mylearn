var <- "let's run R!"
my.function.usefulArray = function(start, end,
isletter = FALSE, isLETTER = FALSE) {

    if (isletter) {
        letters[start:end]
    }

    if (isLETTER) {
        LETTERS[start:end]
    }
}

my.function.fbnq = function(n) {
    a = 1
    b = 1
    while (n) {
        print(a)
        temp = a
        a = a + b
        b = temp
        n = n-1
    }
}

# path = "C:/Users/WHL/Desktop/test/output.txt"
my.function.fromFile = function (path) {
    .fromFile = readLines(path)
}

my.function.toFile = function(path, data) {
    cat(data, file = path)
}

my.function.loopType = function(time) {
    stop = time
    # repeat loop
    repeat {
        print(stop)
        if (stop == 0) {
            break
        }
    }

    stop = time
    # while loop
    while (stop) {
        print(stop)
        stop = stop-1
    }
}

my.function.testCUse = function(time) {
    X = c(1:time)
    print( c( rep(X, time)))
}

my.function.deffWithNaAndNull = function () {
    speak <- "NA mean missing, NULL mean no data there"
    test <-  c(NA, NULL, NULL)
    print( length(test))
    print( test)
}

my.function.dms = function () {
    # change to same
    temp = c(1, 2, TRUE, FALSE, "chino", "saber")
    print(temp)

    tempM = cbind(
        Anima = c("chino", "miko"),
        Comic = c(T, F),
        Game = c("donadona", "game"),
        Novel = c(1, 2)
    )
    print(tempM)
    print( mode(tempM))
}

my.function.dataInfo = function(data, 
show = TRUE,
pMode = TRUE, pLength = TRUE, pName = TRUE, pClass = TRUE) {

    if (show) {
        print('data:')
        print(data)
    }

    # mat save some char, mode return char but class return mat
    if (pMode) {
        print('mode:')
        print(mode(data))
    }

    if (pClass) {
        print('class:')
        print(class(data))
    }

    if (pLength) {
        print('length:')
        print(length(data))
    }

    if (pName) {
        print('name:')
        print(names(data))
    }

}

my.function.useIndex = function() {
    num = 10
    data = 1:num
    # space
    print(data[])
    
    print(data[1])
    index = seq(1, num, by=2)
    print(data[index])

    # ignore
    ignoreIdx = seq(-1, -num, by=-2)
    print(data[ignoreIdx])
    print(data[c(1:2, (num-2):num)])
    print(data[-index])

    # logical
    logicalIdx = c(T, F, T, T, F)
    print(data[logicalIdx])
    print(data < 3)
    print(data[ data<3])
}

my.function.useMatrix = function() {
    # creat , none weill repet
    print( cbind(1:3, 3:1, 1)) #col
    print( rbind(1:3, 3:1, 1)) #row

    # create with c
    tempMatrix = matrix(seq(-6, 5), nrow = 3, ncol = 4, byrow = FALSE)
    # t
    print(t(tempMatrix))

    # size
    print(dim(tempMatrix))
    size = c(row = nrow(tempMatrix), COL = ncol(tempMatrix))
    print(size)

    # put name
    dimnames(tempMatrix) <- list(c("a", "b", "c"), c("chino","yiliya","amiya","k"))
    print(tempMatrix)
    print(dimnames(tempMatrix))

    # index
    print(tempMatrix[,])
    print(tempMatrix[1:2, c(1, 3)])
    print(tempMatrix[-1, -2])

    # the data will keep matrix type with setting 'drop=false'
    print(tempMatrix[1,])
    print(tempMatrix[1, , drop=FALSE])

    #logical
    print(tempMatrix[c(T,F,T),])
    print(tempMatrix[ , tempMatrix[1,]>-1])

    # *
    aMatrix = matrix(c(1, 1, 1, 1, 1, 1, 1, 1, 1), nrow = 3)
    bMatrix = matrix(c(1, 1, 1, 1, 1, 1, 1, 1, 1), nrow = 3)
    print(aMatrix %*% bMatrix)
}

my.function.useArray = function() {
    tempArray = array(1:24, dim = c(4,3,2))
    print(tempArray)

    # info
    print(dim(tempArray))
    dimnames(tempArray) <- list(letters[1:4], LETTERS[5:7], c('inner', 'outter'))
    print(tempArray)
    print(dimnames(tempArray))

    # search also with idx -idx logical
    print(tempArray[tempArray[1,1,]>6,,])
    print(tempArray[-1,1:2,1:2])

    # change size
    dim(tempArray) <- c(4, 6)
    print(tempArray)

    # cover name, never return
    print(dimnames(tempArray))
    dim(tempArray) <- c(4, 3, 2)
    print(dimnames(tempArray))
}

my.function.useList = function() {
    # create empty
    emptyList <- list()
    # the copy, not source
    hasList = list(
        "VEC" = c(1:2, 3:4), 
        "MAT" = matrix(LETTERS[1:6], nrow = 2))
    print(hasList)

    # info
    print(length(hasList))
    print(mode(hasList))

    # index with: space idx -idx logical char
    print(-1)
    print( hasList[T])
    print( hasList["MAT"])

    # ref
    print( hasList[[1]])
    print( hasList$MAT)
    print( hasList$M) # also right

    # add
    hasList$More = letters[10:12]
    print( hasList)
    # with fix NULL
    hasList[[5]] = LETTERS[5:8]
    print( hasList)

    # link
    emptyList$UnEmpty = c('donadona', 'amd yes')
    print( c(hasList, emptyList))
}

my.function.listFunc = function() {
    nExtremes = rpois(100, 3)
    print(nExtremes[1:5])

    exFunc = function(N) round(rweibull(N, shape= 5, scale = 1000))
    extremesValues = lapply(nExtremes, exFunc)
    print(extremesValues[1:5])
}

my.function.DataFrameUse = function() {
    # init
    comic <- data.frame(
        Name = c("chino", "cocoa", "yiliya", "antna", "poruno"),
        From = c("rabbit", "rabbit", "fate", "donadona", "donadona"),
        tag = c("coffee", "sister", "magic", "hello CQ", "driver"),
        id = c(3, 1, 44, 23, 12)
    )

    # info
    length(comic)
    names(comic)

    # col
    print(comic$tag)
    pairs(comic)
}

# checkData = c(APath = rep( c(1, 2, 3), c(2, 3, 1)), BPath = c(4, 5, 6))
# my.function.dataInfo(checkData)

test = "A"
x = switch(test, 
A = "A Word", B = "Boots")
print(x)

# 9-14
# print(getwd())

# print(solve(m))
# print(switch("tet", tet="?"))