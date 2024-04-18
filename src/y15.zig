const std = @import("std");

pub fn solutionD1P1(input: []u8) i32 {
    var floor: i32 = 0;
    for (input) |char| {
        if (char == '(') {
            floor += 1;
        } else {
            floor -= 1;
        }
    }

    return floor;
}

pub fn solutionD1P2(input: []u8) i32 {
    var floor: i32 = 0;
    var pos: i32 = 0;
    for (input) |char| {
        pos += 1;
        if (char == '(') {
            floor += 1;
        } else {
            floor -= 1;
        }
        if (floor == -1) break;
    }

    return pos;
}

pub fn solutionD2P1(input: []u8) i32 {
    var total: i32 = 0;
    var linesIter = std.mem.split(u8, input, "\n");

    while (linesIter.next()) |line| {
        // to handle last empty string
        if (line.len == 0) continue;

        var dimensions: [3]i32 = undefined;
        var idx: usize = 0;
        var lineIter = std.mem.split(u8, line, "x");
        while (lineIter.next()) |dimension| {
            if (std.fmt.parseInt(i32, dimension, 10)) |parsedInt| {
                dimensions[idx] = parsedInt;
            } else |err| {
                std.debug.print("err: {!}\n", .{err});
                break;
            }
            idx += 1;
        }
        idx = 0;

        const s1 = dimensions[0] * dimensions[1];
        const s2 = dimensions[1] * dimensions[2];
        const s3 = dimensions[2] * dimensions[0];

        const boxArea = 2 * s1 + 2 * s2 + 2 * s3;
        const extra = @min(@min(s1, s2), s3);

        total += boxArea + extra;
    }

    return total;
}

pub fn solutionD2P2(input: []u8) i32 {
    var total: i32 = 0;
    var linesIter = std.mem.split(u8, input, "\n");

    while (linesIter.next()) |line| {
        // to handle last empty string
        if (line.len == 0) continue;

        var dimensions: [3]i32 = undefined;
        var idx: usize = 0;
        var lineIter = std.mem.split(u8, line, "x");
        while (lineIter.next()) |dimension| {
            if (std.fmt.parseInt(i32, dimension, 10)) |parsedInt| {
                dimensions[idx] = parsedInt;
            } else |err| {
                std.debug.print("err: {!}\n", .{err});
                break;
            }
            idx += 1;
        }
        idx = 0;

        const smallestSide: i32 = @min(@min(dimensions[0], dimensions[1]), dimensions[2]);
        var smallerSide: i32 = 0;
        if (smallestSide == dimensions[0]) {
            smallerSide = @min(dimensions[1], dimensions[2]);
        } else if (smallestSide == dimensions[1]) {
            smallerSide = @min(dimensions[0], dimensions[2]);
        } else {
            smallerSide = @min(dimensions[0], dimensions[1]);
        }

        const perimeter = 2 * smallestSide + 2 * smallerSide;
        const bow = dimensions[0] * dimensions[1] * dimensions[2];

        total += perimeter + bow;
    }

    return total;
}

pub fn solutionD3P1(input: []u8) !i32 {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var history = std.AutoHashMap([2]i32, bool).init(allocator);
    defer history.deinit();
    var posX: i32 = 0;
    var posY: i32 = 0;

    try history.put([_]i32{ posX, posY }, true);
    var visitedHouses: i32 = 1;

    for (input) |char| {
        if (char == '^') {
            posY += 1;
        } else if (char == '>') {
            posX += 1;
        } else if (char == 'v') {
            posY -= 1;
        } else if (char == '<') {
            posX -= 1;
        }

        const key = [_]i32{ posX, posY };
        const exists = history.contains(key);
        if (!exists) {
            visitedHouses += 1;
        }

        try history.put(key, true);
    }

    return visitedHouses;
}

pub fn solutionD3P2(input: []u8) !i32 {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var historySanta = std.AutoHashMap([2]i32, bool).init(allocator);
    defer historySanta.deinit();
    var santaX: i32 = 0;
    var santaY: i32 = 0;
    try historySanta.put([_]i32{ santaX, santaY }, true);

    var historyRobo = std.AutoHashMap([2]i32, bool).init(allocator);
    defer historyRobo.deinit();
    var roboX: i32 = 0;
    var roboY: i32 = 0;
    try historyRobo.put([_]i32{ roboX, roboY }, true);

    var visitedHouses: i32 = 1;
    var idx: i32 = 0;

    for (input) |char| {
        var command = [_]i32{ 0, 0 };
        if (char == '^') {
            command[1] += 1;
        } else if (char == '>') {
            command[0] += 1;
        } else if (char == 'v') {
            command[1] -= 1;
        } else if (char == '<') {
            command[0] -= 1;
        }

        const isSanta = @mod(idx, 2) == 0;
        var key: [2]i32 = undefined;
        if (isSanta) {
            santaX += command[0];
            santaY += command[1];
            key = [_]i32{ santaX, santaY };
        } else {
            roboX += command[0];
            roboY += command[1];
            key = [_]i32{ roboX, roboY };
        }

        const exists = historySanta.contains(key) or historyRobo.contains(key);
        if (!exists) visitedHouses += 1;

        if (isSanta) try historySanta.put(key, true) else try historyRobo.put(key, true);

        idx += 1;
    }

    return visitedHouses;
}

pub fn solutionD4P1(input: []u8) !i32 {
    _ = input;
    const sol: i32 = 0;

    return sol;
}

pub fn solutionD5P1(input: []u8) !i32 {
    _ = input;
    const sol: i32 = 0;

    return sol;
}
