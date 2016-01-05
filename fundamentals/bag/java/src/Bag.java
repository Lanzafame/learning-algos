import java.util.Iterator;

/**
 * Created by lanzafame on 22/04/15.
 */

/**
 * Bag()
 *
 * Creates an empty bag
 */
public class Bag<Item> implements Iterable<Item> {

    private Node first;     // first node in list
    private int N;          // number of items

    private class Node {
        Item item;
        Node next;
    }

    /**
     * boolean isEmpty()
     *
     * Return a boolean determined by whether the bag is empty or not
     * @return b
     */
    public boolean isEmpty() { return first == null; }

    /**
     * int size()
     *
     * Return the number of items contained in the bag.
     * @return s
     */
    public int size() { return N; }

    /**
     * void add(Item item)
     *
     * Add an item to the Bag
     */
    public void add(Item item) {
        // same as push() in Stack
        Node oldfirst = first;
        first = new Node();
        first.item = item;
        first.next = oldfirst;
    }

    public Iterator<Item> iterator() { return new ListIterator();  }

    private class ListIterator implements Iterator<Item> {
        private Node current = first;

        public boolean hasNext() { return current != null;  }

        public void remove() {}

        public Item next() {
            Item item = current.item;
            current = current.next;
            return item;
        }
    }

}
